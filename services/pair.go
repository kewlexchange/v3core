package services

import (
	"core/models"
	"core/utils"
	"core/workers"
	exchange "core/workers/exchange"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

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
		Test:         currency.Contract.Hex(),
		Currency:     currency,
		TradingPairs: []models.Pair{*p},
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

	// Dosya adı oluştur (ör: output/prefix-<timestamp>.json)
	filename := filepath.Join(outputDir, fmt.Sprintf("%s-%s.json", prefix, "-"))

	// Dosyaya yaz
	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		return "", fmt.Errorf("write file error: %w", err)
	}

	return filename, nil
}

func (s *PairService) ParseTokens(pairs []models.Pair) ([]models.Pair, []models.Asset, error) {
	if len(pairs) == 0 {
		return pairs, []models.Asset{}, nil
	}

	stablePairAddress := pairs[0].Exchange.StablePair
	var stablePair models.Pair
	for _, pair := range pairs {
		if pair.Pair == stablePairAddress.Hex() {
			stablePair = pair
			break
		}
	}

	usdAddress := ""
	nativeAddress := ""

	if stablePair.Base == stablePair.Exchange.NativeToken.Hex() {
		usdAddress = stablePair.Quote
		nativeAddress = stablePair.Base
	} else {
		usdAddress = stablePair.Base
		nativeAddress = stablePair.Quote
	}

	one := decimal.NewFromInt(1)

	if stablePair.Base == nativeAddress {
		// NATIVE / USD
		stablePair.BasePriceNative = &one
		stablePair.QuotePriceUSD = &one

		stablePair.BasePriceUSD = stablePair.BasePrice
		inv := one.Div(*stablePair.BasePrice)
		stablePair.QuotePriceNative = &inv
	} else if stablePair.Base == usdAddress {
		// USD / NATIVE
		stablePair.QuotePriceNative = &one
		stablePair.BasePriceUSD = &one

		stablePair.QuotePriceUSD = stablePair.QuotePrice
		inv := one.Div(*stablePair.QuotePrice)
		stablePair.BasePriceNative = &inv
	}

	assetList := []models.Asset{}

	addPairToAsset(&assetList, stablePair.BaseCurrency, &stablePair)
	addPairToAsset(&assetList, stablePair.QuoteCurrency, &stablePair)

	for id, pair := range pairs {
		if pair.Pair == stablePair.Pair {
			continue
		}

		isBaseNative := pair.Base == nativeAddress || pair.Base == "0x721EF6871f1c4Efe730Dce047D40D1743B886946"
		isQuoteNative := pair.Quote == nativeAddress || pair.Quote == "0x721EF6871f1c4Efe730Dce047D40D1743B886946"
		isBaseStable := pair.Base == stablePair.Base
		isQuoteStable := pair.Quote == stablePair.Quote

		// 1 CHZ = X USD
		var chzUSD decimal.Decimal

		if stablePair.Base == nativeAddress {
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

		} else {
			fmt.Println("Diffrent", pair.Pair)
		}

	}

	for i := range pairs {
		p := &pairs[i]
		if !p.IsEnabled || p.Pair == stablePair.Pair {
			continue
		}

		addPairToAsset(&assetList, p.BaseCurrency, p)
		addPairToAsset(&assetList, p.QuoteCurrency, p)
	}

	return pairs, assetList, nil
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

	pairs, assets, _ := s.ParseTokens(allPairs)
	// Şimdi tüm pairs tek dosyaya yaz
	s.SaveJSONToFile("output", "all_assets", assets)
	s.SaveJSONToFile("output", "all_exchanges", pairs)

	//exchange.DetectArbitrageOnPairs(allPairs)
}
