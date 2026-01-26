package dexv2

import (
	"context"
	"core/constants"
	"core/models"
	"core/services"
	"core/utils"
	"core/workers/exchange/dexv2/contracts/kewl"
	"core/workers/exchange/dexv2/contracts/multicall3"
	"core/workers/exchange/dexv2/contracts/v2Factory"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type DexV2Fetcher struct{}

func (d *DexV2Fetcher) FetchPairs(ex models.Exchange) ([]models.Pair, error) {

	evmClient, err := services.GetEVMClient(*ex.ChainID, *ex.RPC)
	if err != nil {
		return nil, err
	}
	defer evmClient.Close()

	factory, err := v2Factory.NewV2Factory(*ex.Factory, evmClient)
	if err != nil {
		return nil, err
	}

	factoryABI, err := abi.JSON(strings.NewReader(v2Factory.V2FactoryABI))
	if err != nil {
		return nil, err
	}

	multicall, err := multicall3.NewMulticall3(*ex.Multicall3, evmClient)
	if err != nil {
		return nil, err
	}

	length, err := factory.AllPairsLength(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	calls := make([]multicall3.Multicall3Call3, length.Int64())
	for i := int64(0); i < length.Int64(); i++ {
		data, err := factoryABI.Pack("allPairs", big.NewInt(i))
		if err != nil {
			return nil, err
		}
		calls[i] = multicall3.Multicall3Call3{
			Target:   *ex.Factory,
			CallData: data,
		}
	}

	multicallRaw := &multicall3.Multicall3Raw{Contract: multicall}
	callOpts := &bind.CallOpts{Context: context.Background()}

	var rawResult []interface{}
	err = multicallRaw.Call(callOpts, &rawResult, "aggregate3", calls)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(rawResult[0])
	if err != nil {
		return nil, fmt.Errorf("json marshal failed: %w", err)
	}

	var results []multicall3.Multicall3Result
	if err := json.Unmarshal(jsonBytes, &results); err != nil {
		return nil, fmt.Errorf("json unmarshal failed: %w", err)
	}
	pairs := make([]models.Pair, 0, length.Int64())
	for _, res := range results {
		if !res.Success {
			continue
		}

		out, err := factoryABI.Unpack("allPairs", res.ReturnData)
		if err != nil {
			log.Printf("Failed to unpack pair address: %v", err)
			continue
		}

		pairAddress, ok := out[0].(common.Address)
		if !ok {
			log.Printf("Failed to cast unpacked output to common.Address, got %T", out[0])
			continue
		}

		pairs = append(pairs, models.Pair{
			Pair:     pairAddress.Hex(),
			Exchange: ex,
		})

	}

	log.Printf("[DEX Fetcher] fetched %d pairs for %s", len(pairs), ex.Name)

	pairs, err = d.FetchReserves(pairs)
	return pairs, err
}

func (d *DexV2Fetcher) CalculatePrices(reserveBase, reserveQuote, decimalsBase, decimalsQuote *big.Int) (priceBase, priceQuote decimal.Decimal, err error) {
	decimalsBaseInt64 := decimalsBase.Int64()
	decimalsQuoteInt64 := decimalsQuote.Int64()

	baseDivisor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(decimalsBaseInt64), nil))
	quoteDivisor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(decimalsQuoteInt64), nil))

	fReserveBase := new(big.Float).Quo(new(big.Float).SetInt(reserveBase), baseDivisor)
	fReserveQuote := new(big.Float).Quo(new(big.Float).SetInt(reserveQuote), quoteDivisor)

	zero := big.NewFloat(0)

	if fReserveBase.Cmp(zero) == 0 || fReserveQuote.Cmp(zero) == 0 {
		return decimal.Zero, decimal.Zero, fmt.Errorf("reserve base veya quote sıfır")
	}

	priceBaseFloat := new(big.Float).Quo(fReserveQuote, fReserveBase)
	priceQuoteFloat := new(big.Float).Quo(fReserveBase, fReserveQuote)

	// big.Float → string
	priceBaseStr := priceBaseFloat.Text('f', 18) // 18 ondalık hassasiyet
	priceQuoteStr := priceQuoteFloat.Text('f', 18)

	// string → decimal.Decimal
	priceBase, err = decimal.NewFromString(priceBaseStr)
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	priceQuote, err = decimal.NewFromString(priceQuoteStr)
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}

	return priceBase, priceQuote, nil
}

func (d *DexV2Fetcher) FetchReserves(pairs []models.Pair) ([]models.Pair, error) {

	kewlInfo, _ := constants.GetExchangeByName("KEWL")
	evmClient, err := services.GetEVMClient(*kewlInfo.ChainID, *kewlInfo.RPC)
	if err != nil {
		return nil, err
	}
	defer evmClient.Close()

	kewlSwap, err := kewl.NewKewl(*kewlInfo.Factory, evmClient)
	if err != nil {
		return nil, err
	}

	var pairlist []common.Address
	for _, pair := range pairs {
		pairlist = append(pairlist, common.HexToAddress(pair.Pair))
	}

	tradePairs, err := kewlSwap.GetReservesByPairAddresses(&bind.CallOpts{}, pairlist)
	if err != nil {
		log.Fatalf("Error fetching reserves: %v", err)
	}

	// tradePairs artık []kewlSwap.TradePairInfo tipinde
	var minLiquidity = big.NewInt(1000) // 1000 Wei minimum likidite
	for i, pair := range tradePairs {
		pairAddress := common.HexToAddress(pairs[i].Pair)
		if pairAddress.Hex() == pair.Pair.Hex() {
			pairs[i].Base = pair.Token0.Hex()
			pairs[i].Quote = pair.Token1.Hex()
			pairs[i].BaseReserve = pair.Reserve0
			pairs[i].QuoteReserve = pair.Reserve1

			pairs[i].BaseReserveStr = utils.StringPtr(utils.FormatUnits(pair.Reserve0, pair.Token0Decimals))
			pairs[i].QuoteReserveStr = utils.StringPtr(utils.FormatUnits(pair.Reserve1, pair.Token1Decimals))
			pairs[i].BaseDecimals = pair.Token0Decimals
			pairs[i].QuoteDecimals = pair.Token1Decimals

			pairs[i].BaseCurrency.Contract = &pair.Token0
			pairs[i].BaseCurrency.Decimals = pair.Token0Decimals

			pairs[i].QuoteCurrency.Contract = &pair.Token1
			pairs[i].QuoteCurrency.Decimals = pair.Token1Decimals

			var priceBase decimal.Decimal
			var priceQuote decimal.Decimal
			priceBase, priceQuote, err = d.CalculatePrices(pair.Reserve0, pair.Reserve1, pair.Token0Decimals, pair.Token1Decimals)

			pairs[i].BasePrice = &priceBase
			pairs[i].QuotePrice = &priceQuote

			isEnabled := pair.Reserve0.Cmp(minLiquidity) >= 0 && pair.Reserve1.Cmp(minLiquidity) >= 0
			pairs[i].IsEnabled = isEnabled
		}
	}

	return pairs, nil
}

/*
abigen --abi=workers/exchange/dexv2/abis/v2Factory.abi --pkg=v2Factory --out=workers/exchange/dexv2/contracts/v2Factory/v2Factory.go
abigen --abi=workers/exchange/dexv2/abis/multicall3.abi --pkg=multicall3  --out=workers/exchange/dexv2/contracts/multicall3/multicall3.go
abigen --abi=workers/exchange/dexv2/abis/v2Pair.abi --pkg=v2Pair --out=workers/exchange/dexv2/contracts/v2Pair/v2Pair.go
abigen --abi=workers/exchange/dexv2/abis/kewl.abi --pkg=kewl --out=workers/exchange/dexv2/contracts/kewl/kewl.go

*/
