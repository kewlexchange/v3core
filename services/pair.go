package services

import (
	"core/models"
	"core/workers"
	exchange "core/workers/exchange"
	"core/workers/exchange/dexv2/scanner"
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

func addPairToAsset(assetList *[]models.Asset, currency models.Currency, p *models.TradingPair) {
	for i := range *assetList {
		if (*assetList)[i].Currency.Contract.Hex() == currency.Contract.Hex() {
			(*assetList)[i].TradingPairs = append((*assetList)[i].TradingPairs, *p)
			return
		}
	}
	*assetList = append(*assetList, models.Asset{
		ContractAddress: currency.Contract.Hex(),
		Currency:        currency,
		TradingPairs:    []models.TradingPair{*p},
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

func (s *PairService) PriceNativeOf(asset common.Address, p *models.TradingPair) *decimal.Decimal {
	if p.Base == asset {
		return p.BasePriceNative
	}
	if p.Quote == asset {
		return p.QuotePriceNative
	}
	return nil
}

func (s *PairService) PriceUSDOf(asset common.Address, p *models.TradingPair) *decimal.Decimal {
	if p.Base == asset {
		return p.BasePriceUSD
	}
	if p.Quote == asset {
		return p.QuotePriceUSD
	}
	return nil
}

func (s *PairService) FindPair(assets []models.Asset, currency models.Currency, stableOrNativeToken common.Address) *models.TradingPair {

	token := currency.Contract

	var bestPair *models.TradingPair
	var bestLiquidity *big.Int
	for _, asset := range assets {
		if asset.Currency.Contract.Hex() == currency.Contract.Hex() {

			for _, pair := range asset.TradingPairs {

				if (pair.Base == stableOrNativeToken && pair.Quote == *token) ||
					(pair.Base == *token && pair.Quote == stableOrNativeToken) {

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

func (s *PairService) FetchPairsConcurrent(exchanges []models.Exchange) {
	var wg sync.WaitGroup
	mu := &sync.Mutex{}                // pairs slice erişimi için mutex
	allPairs := []models.TradingPair{} // veya pairs'in tipi neyse onu kullan

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

	s.SaveJSONToFile("output", "all_exchanges", allPairs)
}

func (s *PairService) ScanPairs(params []models.ScanParams) {
	allPairs := []models.TradingPair{} // veya pairs'in tipi neyse onu kullan

	for _, param := range params {
		fmt.Println("Scanner", param.Token)

		pairs, err := s.fetcher.FetchReserves(param.Pairs)
		if err != nil {
			log.Printf("error fetching for %s: %v", param.Token, err)
			return
		}
		allPairs = append(allPairs, pairs...)
		scanner.FlashSwap(param, pairs)
		s.SaveJSONToFile("output", param.Token, pairs)
	}
}
