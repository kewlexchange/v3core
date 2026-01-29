package dexv2

import (
	"context"
	"core/constants"
	"core/models"
	"core/services"
	coreTypes "core/types"
	"core/utils"
	"core/workers/exchange/dexv2/contracts/flash"
	"core/workers/exchange/dexv2/contracts/kewl"
	"core/workers/exchange/dexv2/contracts/multicall3"
	"core/workers/exchange/dexv2/contracts/v2Factory"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
)

type DexV2Fetcher struct{}

func (d *DexV2Fetcher) FetchPairs(ex models.Exchange) ([]models.TradingPair, error) {

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
	pairs := make([]models.TradingPair, 0, length.Int64())
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

		pairs = append(pairs, models.TradingPair{
			Pair:     pairAddress,
			Exchange: ex,
		})

	}

	log.Printf("[DEX Fetcher] fetched %d pairs for %s", len(pairs), ex.Name)

	pairs, err = d.FetchReserves(*ex.ChainID, pairs)
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

func (d *DexV2Fetcher) FetchReserves(chainId models.ChainID, pairs []models.TradingPair) ([]models.TradingPair, error) {

	fmt.Println("FetchReserves", chainId)
	kewlInfo, _ := constants.GetExchangeByName("KEWL", chainId)
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
		pairlist = append(pairlist, pair.Pair)
	}

	tradePairs, err := kewlSwap.GetReservesByPairAddresses(&bind.CallOpts{}, pairlist)
	if err != nil {
		log.Fatalf("Error fetching reserves: %v", err)
	}

	// tradePairs artık []kewlSwap.TradePairInfo tipinde
	var minLiquidity = big.NewInt(1000) // 1000 Wei minimum likidite
	for i, pair := range tradePairs {
		pairAddress := pairs[i].Pair
		if pairAddress.Hex() == pair.Pair.Hex() {
			base := pair.Token0.Hex()
			quote := pair.Token1.Hex()

			pairs[i].Base = &base
			pairs[i].Quote = &quote

			pairs[i].BaseReserve = pair.Reserve0
			pairs[i].QuoteReserve = pair.Reserve1

			pairs[i].BaseReserveStr = utils.StringPtr(utils.FormatUnits(pair.Reserve0, pair.Token0Decimals))
			pairs[i].QuoteReserveStr = utils.StringPtr(utils.FormatUnits(pair.Reserve1, pair.Token1Decimals))
			pairs[i].BaseDecimals = pair.Token0Decimals
			pairs[i].QuoteDecimals = pair.Token1Decimals

			pairs[i].BaseCurrency.ChainID = kewlInfo.ChainID
			pairs[i].BaseCurrency.Contract = &pair.Token0
			pairs[i].BaseCurrency.Decimals = pair.Token0Decimals

			pairs[i].QuoteCurrency.Contract = &pair.Token1
			pairs[i].QuoteCurrency.Decimals = pair.Token1Decimals
			pairs[i].QuoteCurrency.ChainID = kewlInfo.ChainID

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

func (d *DexV2Fetcher) ExecuteSwap(chainId models.ChainID, params coreTypes.ArbResult) error {

	if !params.Exists {
		return nil
	}

	if chainId != constants.ChilizChainId {
		fmt.Println("BSC/AVAX")
		return nil
	}

	flashContract := constants.FlashContractMap[models.ChainID(chainId)]
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	if privateKeyHex == "" {
		fmt.Println("PRIVATE_KEY env variable is not set")
		return fmt.Errorf("Invalid Private Key")
	}

	evmClient := services.Clients[chainId]

	flashSwap, err := flash.NewFlash(flashContract, evmClient)
	if err != nil {
		fmt.Println("Err1", err)
		return err
	}

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		fmt.Println("Err2", err)
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	from := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := evmClient.PendingNonceAt(context.Background(), from)
	if err != nil {
		fmt.Println("Err3", err)

		return err
	}

	gasPrice, err := evmClient.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	chainID := big.NewInt(chainId.Int64())
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return err
	}

	auth.From = from
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)

	auth.GasLimit = 0 // uint64(2_600_000)
	auth.GasPrice = gasPrice

	flashParams := flash.SwapFlashParams{

		DX:     params.Dx,
		Profit: params.Profit,
		Mid:    params.Mid,
		Out:    params.Out,
		Borrow: params.Borrow,
		Output: params.Repay,
		Path:   params.Path,
	}

	if !tryLockFlash(flashParams) {
		fmt.Println("Daha once islendi...")
		return nil
	}
	defer unlockFlash(flashParams)
	/*
		// ethclient, rpcClient ile sarmalanmış
		ethClient := ethclient.NewClient(evmClient.Client())

		parsedABI, err := abi.JSON(strings.NewReader(flash.FlashABI))
		if err != nil {
			log.Fatalf("Failed to parse ABI: %v", err)
		}

		// 2. handleFlash fonksiyonunu bul
		method := parsedABI.Methods["handleFlash"]

		// 4. Parametreleri encode et
		encodedParams, err := method.Inputs.Pack(flashParams)
		if err != nil {
			log.Fatalf("Failed to pack params: %v", err)
		}
		data := append(method.ID, encodedParams...)

		msg := ethereum.CallMsg{
			To:   &common.Address{},
			Data: data,
		}


		result, err := ethClient.CallContract(context.Background(), msg, nil)
		if err != nil {
			log.Fatalf("CallContract error: %v", err)
		}

		log.Printf("Result: %x", result)

		fmt.Println("eth_call result:", result)
	*/
	fmt.Println("DX", params.Dx)
	fmt.Println("Mid", params.Mid)
	fmt.Println("Out", params.Out)
	fmt.Println("Borrow", params.Borrow)
	fmt.Println("Repay", params.Repay)
	fmt.Println("Pair0", params.Path[0].Hex())
	fmt.Println("Pair1", params.Path[1].Hex())

	tx, err := flashSwap.HandleSwap(auth, flashParams)
	if err != nil {
		fmt.Println("Coder5", err)

		return err
	}

	fmt.Println("Flash TX:", tx.Hash().Hex())
	return nil

	//flashSwap.HandleFlash()
}

func (d *DexV2Fetcher) ExecuteSwapAll(chainId models.ChainID, params []coreTypes.ArbResult) error {
	if len(params) == 0 {
		fmt.Println("EXEC_ALL_ZERO_PAIR")
		return nil
	}

	flashContract := constants.FlashContractMap[models.ChainID(chainId)]
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	if privateKeyHex == "" {
		fmt.Println("PRIVATE_KEY env variable is not set")
		return fmt.Errorf("Invalid Private Key")
	}

	evmClient := services.Clients[chainId]

	flashSwap, err := flash.NewFlash(flashContract, evmClient)
	if err != nil {
		fmt.Println("Err1", err)
		return err
	}

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		fmt.Println("Err2", err)
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	from := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := evmClient.PendingNonceAt(context.Background(), from)
	if err != nil {
		fmt.Println("Err3", err)

		return err
	}

	gasPrice, err := evmClient.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	chainID := big.NewInt(chainId.Int64())
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return err
	}

	auth.From = from
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)

	auth.GasLimit = 0 // uint64(2_600_000)
	auth.GasPrice = gasPrice

	allSwapParams := []flash.SwapFlashParams{}

	totalProfit := big.NewInt(0)
	for _, param := range params {
		if !param.Exists {
			continue
		}
		flashParams := flash.SwapFlashParams{
			DX:     param.Dx,
			Profit: param.Profit,
			Mid:    param.Mid,
			Out:    param.Out,
			Borrow: param.Borrow,
			Output: param.Repay,
			Path:   param.Path,
		}
		allSwapParams = append(allSwapParams, flashParams)
		if param.Profit != nil {
			totalProfit.Add(totalProfit, param.Profit)
		}
	}

	fmt.Println("TOTAL PROFIT", utils.FormatUnits(totalProfit, big.NewInt(18)))

	estimateAuth := *auth      // 👈 signer, chainID, her şey kopyalanır
	estimateAuth.NoSend = true // sadece gönderme kapalı

	input, err := flashSwap.FlashTransactor.HandleSwapEx(
		&estimateAuth,
		allSwapParams,
	)
	if err != nil {
		fmt.Println("ERR", err)
		return err
	}

	estimatedGas := input.Gas()

	estimatedFee := new(big.Int).Mul(
		big.NewInt(int64(estimatedGas)),
		gasPrice,
	)

	fmt.Println("Estimated Gas: ", estimatedGas, estimatedFee, totalProfit)

	if estimatedFee.Cmp(totalProfit) > 0 {
		fmt.Println("❌ Gas fee profitten büyük, işlem iptal")
		return nil
	}

	if chainId != constants.ChilizChainId {
		fmt.Println("BSC AVAX")
		return nil
	}

	tx, err := flashSwap.HandleSwapEx(auth, allSwapParams)
	if err != nil {
		fmt.Println("HandleSwapEx", err)

		return err
	}

	fmt.Println("Tx sent:", tx.Hash().Hex())

	// ⏳ confirm bekle
	receipt, err := bind.WaitMined(context.Background(), evmClient, tx)
	if err != nil {
		fmt.Println("WaitMined error:", err)
		return err
	}

	if receipt.Status == types.ReceiptStatusFailed {
		fmt.Println("❌ Tx reverted")
		return fmt.Errorf("transaction reverted")
	}

	fmt.Println("✅ Tx confirmed")
	fmt.Println("Gas used:", receipt.GasUsed)

	fmt.Println("Flash TX:", tx.Hash().Hex())
	return nil
}

/*
abigen --abi=workers/exchange/dexv2/abis/v2Factory.abi --pkg=v2Factory --out=workers/exchange/dexv2/contracts/v2Factory/v2Factory.go
abigen --abi=workers/exchange/dexv2/abis/multicall3.abi --pkg=multicall3  --out=workers/exchange/dexv2/contracts/multicall3/multicall3.go
abigen --abi=workers/exchange/dexv2/abis/v2Pair.abi --pkg=v2Pair --out=workers/exchange/dexv2/contracts/v2Pair/v2Pair.go
abigen --abi=workers/exchange/dexv2/abis/kewl.abi --pkg=kewl --out=workers/exchange/dexv2/contracts/kewl/kewl.go
abigen --abi=workers/exchange/dexv2/abis/flash.abi --pkg=flash --out=workers/exchange/dexv2/contracts/flash/flash.go

*/
