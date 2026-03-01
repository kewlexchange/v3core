package cex

import (
	"core/models"
	coreTypes "core/types"
	"encoding/json"
	"fmt"
	"log"
)

type CexFetcher struct {
	client any
}

func NewCexFetcher(client any) *CexFetcher {
	return &CexFetcher{client: client}
}

func (c *CexFetcher) FetchPairs(exchange models.Exchange) ([]models.TradingPair, error) {
	log.Printf("[CEX Fetcher] Fetching pairs from %s ...", exchange.Name)

	// CCXT-Go async LoadMarkets çağrısı
	ex, ok := c.client.(interface {
		LoadMarkets(...interface{}) <-chan interface{}
	})
	if !ok {
		return nil, fmt.Errorf("client does not support LoadMarkets")
	}

	ch := ex.LoadMarkets()
	resp := <-ch

	// Gelen response doğrudan root-level market map
	raw, ok := resp.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid CCXT response (not a map)")
	}

	// Debug istersen:
	pretty, _ := json.MarshalIndent(raw, "", "  ")
	fmt.Println(string(pretty))

	pairs := []models.TradingPair{}

	// ÖNEMLİ → raw doğrudan "symbol → market"
	for symbol, entry := range raw {
		market, ok := entry.(map[string]interface{})
		if !ok {
			continue
		}

		base, _ := market["base"].(string)
		quote, _ := market["quote"].(string)

		// Debug:
		// fmt.Println("Market:", symbol, "Base:", base, "Quote:", quote)

		if base == "" || quote == "" {
			continue
		}

		pair := models.TradingPair{
			ExchangeID: exchange.ID,
			Symbol:     symbol,
			Base:       &base,
			Quote:      &quote,
		}

		fmt.Println("base", exchange.Name, base, quote)
		pairs = append(pairs, pair)
	}

	log.Printf("[CEX] %s → %d pairs fetched", exchange.Name, len(pairs))
	return pairs, nil
}

func (c *CexFetcher) FetchReserves(chainId models.ChainID, pairs []models.TradingPair) ([]models.TradingPair, error) {
	return nil, nil
}

func (c *CexFetcher) ExecuteSwap(chainId models.ChainID, params coreTypes.ArbResult) error {

	return nil
}

func (c *CexFetcher) ExecuteSwapAll(chainId models.ChainID, params []coreTypes.ArbResult) error {
	return nil
}

func (c *CexFetcher) FetchCycle(chainId models.ChainID, params models.Cycle) (models.Cycle, error) {
	return models.Cycle{}, nil
}

func (c *CexFetcher) FetchBalances() (map[string]float64, error) {
	return nil, nil
}

func (c *CexFetcher) FetchBalance(symbol string) (float64, error) {
	return 0, nil
}

func (c *CexFetcher) CreateMarketBuyOrder(symbol string, amount float64) (string, error) {
	return "", nil

}
func (c *CexFetcher) CreateLimitBuyOrder(symbol string, amount float64, price float64) (string, error) {
	return "", nil
}

func (c *CexFetcher) CreateMarketSellOrder(symbol string, amount float64) (string, error) {
	return "", nil
}

func (c *CexFetcher) CreateLimitSellOrder(symbol string, amount float64, price float64) (string, error) {
	return "", nil
}

func (c *CexFetcher) CancelOrder(id string) error {
	return nil
}

func (c *CexFetcher) CancelAllOrders() error {
	return nil
}

func (c *CexFetcher) CancelOrders(ids []string) error {
	return nil
}
