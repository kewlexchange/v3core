package services

import (
	"core/models"
	"core/utils"
	"core/workers"
	exchange "core/workers/exchange"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type PairService struct {
	pool    *workers.WorkerPool
	fetcher exchange.PairFetcher
}

func NewPairService(pool *workers.WorkerPool, fetcher exchange.PairFetcher) *PairService {
	return &PairService{
		pool:    pool,
		fetcher: fetcher,
	}
}

func addPairToAsset(assetList *[]models.Asset, currency models.Currency, p *models.Pair) {
	for i := range *assetList {
		if (*assetList)[i].Currency.Contract.Hex() == currency.Contract.Hex() {
			(*assetList)[i].TradingPairs = append((*assetList)[i].TradingPairs, *p)
			return
		}
	}
	*assetList = append(*assetList, models.Asset{
		ContractAddress: currency.Contract.Hex(),
		Currency:        currency,
		TradingPairs:    []models.Pair{*p},
	})
}

func (s *PairService) SaveJSONToFile(outputDir, prefix string, data interface{}) (string, error) {
	// JSON'a çevir (indentli)
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("json marshal error: %w", err)
	}

	// Klasör oluştur
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("mkdir error: %w", err)
	}

	filename := filepath.Join(outputDir, fmt.Sprintf("%s-%s.json", prefix, "data"))

	// Dosyaya yaz
	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		return "", fmt.Errorf("write file error: %w", err)
	}

	return filename, nil
}

func (s *PairService) PriceNativeOf(asset common.Address, p *models.Pair) *decimal.Decimal {
	if p.Base == asset.Hex() {
		return p.BasePriceNative
	}
	if p.Quote == asset.Hex() {
		return p.QuotePriceNative
	}
	return nil
}

func (s *PairService) PriceUSDOf(asset common.Address, p *models.Pair) *decimal.Decimal {
	if p.Base == asset.Hex() {
		return p.BasePriceUSD
	}
	if p.Quote == asset.Hex() {
		return p.QuotePriceUSD
	}
	return nil
}

func (s *PairService) FindPair(assets []models.Asset, currency models.Currency, stableOrNativeToken string) *models.Pair {

	token := currency.Contract.Hex()

	var bestPair *models.Pair
	var bestLiquidity *big.Int
	for _, asset := range assets {
		if asset.Currency.Contract.Hex() == currency.Contract.Hex() {

			for _, pair := range asset.TradingPairs {

				if (pair.Base == stableOrNativeToken && pair.Quote == token) ||
					(pair.Base == token && pair.Quote == stableOrNativeToken) {

					currentLiquidity := new(big.Int).Mul(pair.BaseReserve, pair.QuoteReserve)

					if bestPair == nil {
						bestPair = &pair
						bestLiquidity = currentLiquidity
					} else if currentLiquidity.Cmp(bestLiquidity) > 0 {
						bestPair = &pair
						bestLiquidity = currentLiquidity
					}
				}
			}
		}
	}
	return bestPair
}

func (s *PairService) ParseTokens(pairs []models.Pair) ([]models.Pair, []models.Asset, models.Pair, error) {
	if len(pairs) == 0 {
		return pairs, []models.Asset{}, models.Pair{}, nil
	}

	stablePairAddress := pairs[0].Exchange.StablePair
	var stablePair models.Pair
	for _, pair := range pairs {
		if pair.Pair == stablePairAddress.Hex() {
			stablePair = pair
			break
		}
	}

	usdAddress := common.Address{}
	nativeAddress := common.Address{}

	if stablePair.Base == stablePair.Exchange.NativeToken.Hex() {
		usdAddress = *utils.AddressFromHex(stablePair.Quote)
		nativeAddress = *utils.AddressFromHex(stablePair.Base)
	} else {
		usdAddress = *utils.AddressFromHex(stablePair.Base)
		nativeAddress = *utils.AddressFromHex(stablePair.Quote)
	}

	one := decimal.NewFromInt(1)

	if stablePair.Base == nativeAddress.Hex() {
		// NATIVE / USD
		stablePair.BasePriceNative = &one
		stablePair.QuotePriceUSD = &one

		stablePair.BasePriceUSD = stablePair.BasePrice
		inv := one.Div(*stablePair.BasePrice)
		stablePair.QuotePriceNative = &inv
	} else if stablePair.Base == usdAddress.Hex() {
		// USD / NATIVE
		stablePair.QuotePriceNative = &one
		stablePair.BasePriceUSD = &one

		stablePair.QuotePriceUSD = stablePair.QuotePrice
		inv := one.Div(*stablePair.QuotePrice)
		stablePair.BasePriceNative = &inv
	}

	assetList := []models.Asset{}
	unknownPairs := []models.Pair{}

	addPairToAsset(&assetList, stablePair.BaseCurrency, &stablePair)
	addPairToAsset(&assetList, stablePair.QuoteCurrency, &stablePair)

	for id, pair := range pairs {
		if pair.Pair == stablePair.Pair {
			continue
		}

		if !pair.IsEnabled {
			continue
		}

		isBaseNative := pair.Base == nativeAddress.Hex() || pair.Base == "0x721EF6871f1c4Efe730Dce047D40D1743B886946"
		isQuoteNative := pair.Quote == nativeAddress.Hex() || pair.Quote == "0x721EF6871f1c4Efe730Dce047D40D1743B886946"
		isBaseStable := pair.Base == stablePair.Base
		isQuoteStable := pair.Quote == stablePair.Quote

		// 1 CHZ = X USD
		var chzUSD decimal.Decimal

		if stablePair.Base == nativeAddress.Hex() {
			// CHZ / USD
			chzUSD = *stablePair.BasePrice
		} else {
			// USD / CHZ
			chzUSD = decimal.NewFromInt(1).Div(*stablePair.BasePrice)
		}
		if isBaseNative || isQuoteNative || isBaseStable || isQuoteStable {
			// CHZ / TOKEN
			if isBaseNative {
				pairs[id].BasePriceNative = utils.DecimalPtr(decimal.NewFromInt(1))
				pairs[id].QuotePriceNative = pair.QuotePrice

				pairs[id].BasePriceUSD = utils.DecimalPtr(chzUSD)
				v := pair.QuotePrice.Mul(chzUSD)
				pairs[id].QuotePriceUSD = &v
			}

			// TOKEN / CHZ
			if isQuoteNative {
				pairs[id].QuotePriceNative = utils.DecimalPtr(decimal.NewFromInt(1))
				pairs[id].BasePriceNative = pair.BasePrice

				pairs[id].QuotePriceUSD = utils.DecimalPtr(chzUSD)
				v := pair.BasePrice.Mul(chzUSD)
				pairs[id].BasePriceUSD = &v
			}

			// USD / TOKEN
			if isBaseStable {
				pairs[id].BasePriceUSD = utils.DecimalPtr(decimal.NewFromInt(1))
				pairs[id].QuotePriceUSD = pair.QuotePrice

				inv := decimal.NewFromInt(1).Div(chzUSD)
				pairs[id].BasePriceNative = &inv

				v := pair.QuotePrice.Mul(inv)
				pairs[id].QuotePriceNative = &v
			}

			// TOKEN / USD
			if isQuoteStable {
				pairs[id].QuotePriceUSD = utils.DecimalPtr(decimal.NewFromInt(1))
				pairs[id].BasePriceUSD = pair.BasePrice

				inv := decimal.NewFromInt(1).Div(chzUSD)
				pairs[id].QuotePriceNative = &inv

				v := pair.BasePrice.Mul(inv)
				pairs[id].BasePriceNative = &v
			}

			addPairToAsset(&assetList, pairs[id].BaseCurrency, &pairs[id])
			addPairToAsset(&assetList, pairs[id].QuoteCurrency, &pairs[id])

		} else {
			unknownPairs = append(unknownPairs, pair)
		}
	}

	for i := range unknownPairs {
		p := &unknownPairs[i]

		baseNativePair := s.FindPair(assetList, p.BaseCurrency, nativeAddress.Hex())
		quoteNativePair := s.FindPair(assetList, p.QuoteCurrency, nativeAddress.Hex())

		if baseNativePair != nil && quoteNativePair != nil {
			baseNative := s.PriceNativeOf(*p.BaseCurrency.Contract, baseNativePair)
			quoteNative := s.PriceNativeOf(*p.QuoteCurrency.Contract, quoteNativePair)

			if baseNative != nil && quoteNative != nil {
				price := baseNative.Div(*quoteNative)
				p.BasePrice = &price
				inv := decimal.NewFromInt(1).Div(price)
				p.QuotePrice = &inv

				p.BasePriceNative = baseNative
				p.QuotePriceNative = quoteNative

				baseUSD := s.PriceUSDOf(*p.BaseCurrency.Contract, baseNativePair)
				quoteUSD := s.PriceUSDOf(*p.QuoteCurrency.Contract, quoteNativePair)
				if baseUSD != nil && quoteUSD != nil {
					p.BasePriceUSD = baseUSD
					p.QuotePriceUSD = quoteUSD
				}

				addPairToAsset(&assetList, p.BaseCurrency, p)
				addPairToAsset(&assetList, p.QuoteCurrency, p)
				continue
			}
		}

		baseUSDPair := s.FindPair(assetList, p.BaseCurrency, stablePairAddress.Hex())
		quoteUSDPair := s.FindPair(assetList, p.QuoteCurrency, stablePairAddress.Hex())

		if baseUSDPair != nil && quoteUSDPair != nil {
			baseUSD := s.PriceUSDOf(*p.BaseCurrency.Contract, baseUSDPair)
			quoteUSD := s.PriceUSDOf(*p.QuoteCurrency.Contract, quoteUSDPair)

			if baseUSD != nil && quoteUSD != nil {
				price := baseUSD.Div(*quoteUSD)
				p.BasePrice = &price
				inv := decimal.NewFromInt(1).Div(price)
				p.QuotePrice = &inv

				p.BasePriceUSD = baseUSD
				p.QuotePriceUSD = quoteUSD

				baseNative := s.PriceNativeOf(*p.BaseCurrency.Contract, baseUSDPair)
				quoteNative := s.PriceNativeOf(*p.QuoteCurrency.Contract, quoteUSDPair)
				if baseNative != nil && quoteNative != nil {
					p.BasePriceNative = baseNative
					p.QuotePriceNative = quoteNative
				}

				addPairToAsset(&assetList, p.BaseCurrency, p)
				addPairToAsset(&assetList, p.QuoteCurrency, p)
				continue
			}
		}

	}

	return pairs, assetList, stablePair, nil
}

func (s *PairService) CustomPairs(assets []models.Asset, baseToken *common.Address, quoteToken *common.Address) []models.Pair {

	uniq := make(map[string]models.Pair)

	for _, asset := range assets {
		addr := asset.Currency.Contract.Hex()
		if addr != baseToken.Hex() && addr != quoteToken.Hex() {
			continue
		}

		for _, pair := range asset.TradingPairs {
			if (pair.Base == baseToken.Hex() && pair.Quote == quoteToken.Hex()) ||
				(pair.Quote == baseToken.Hex() && pair.Base == quoteToken.Hex()) {

				uniq[pair.Pair] = pair
			}
		}
	}

	out := make([]models.Pair, 0, len(uniq))
	for _, p := range uniq {
		out = append(out, p)
	}

	return out
}

func (s *PairService) FetchPairsConcurrent(exchanges []models.Exchange) {
	var wg sync.WaitGroup
	mu := &sync.Mutex{}         // pairs slice erişimi için mutex
	allPairs := []models.Pair{} // veya pairs'in tipi neyse onu kullan

	for _, ex := range exchanges {
		wg.Add(1)
		exCopy := ex
		s.pool.Submit(func() {
			defer wg.Done()
			pairs, err := s.fetcher.FetchPairs(exCopy)
			if err != nil {
				log.Printf("error fetching for %s: %v", exCopy.Name, err)
				return
			}

			mu.Lock()
			allPairs = append(allPairs, pairs...)
			mu.Unlock()

			log.Printf("service received %d pairs from %s", len(pairs), exCopy.Name)
		})
	}

	wg.Wait() // tüm fetch işlemleri bitene kadar bekle

	pairs, assets, _, _ := s.ParseTokens(allPairs)
	customs := s.CustomPairs(assets, utils.AddressFromHex("0x60f397acbcfb8f4e3234c659a3e10867e6fa6b67"), utils.AddressFromHex("0x677f7e16c7dd57be1d4c8ad1244883214953dc47"))
	// Şimdi tüm pairs tek dosyaya yaz
	s.SaveJSONToFile("output", "customs", customs)
	s.SaveJSONToFile("output", "all_assets", assets)
	s.SaveJSONToFile("output", "all_exchanges", pairs)

	//exchange.DetectArbitrageOnPairs(allPairs)
}
