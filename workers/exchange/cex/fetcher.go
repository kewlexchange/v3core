package cex

import (
	"core/models/db"
	"fmt"
	"log"
)

type CexFetcher struct {
	client any
}

func NewCexFetcher(client any) *CexFetcher {
	return &CexFetcher{client: client}
}

func (c *CexFetcher) FetchPairs(exchange db.Exchange) ([]db.Pair, error) {
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

	pairs := []db.Pair{}

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

		pair := db.Pair{
			ExchangeID: exchange.ID,
			Symbol:     symbol,
			Base:       base,
			Quote:      quote,
		}

		fmt.Println("base", exchange.Name, base, quote)
		pairs = append(pairs, pair)
	}

	log.Printf("[CEX] %s → %d pairs fetched", exchange.Name, len(pairs))
	return pairs, nil
}
