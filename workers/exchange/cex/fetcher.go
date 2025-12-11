package cex

import (
	"core/models/db"
	"fmt"
	"log"

	ccxt "github.com/ccxt/ccxt/go/v4"
)

type CexFetcher struct {
	client ccxt.IExchange
}

func (c *CexFetcher) FetchPairs(exchange db.Exchange) ([]db.Pair, error) {
	log.Printf("[CEX Fetcher] Fetching pairs from %s ...", exchange.Name)

	// CCXT-Go async call
	ch := c.client.LoadMarkets()

	resp := <-ch // kanal cevabı
	raw, ok := resp.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid CCXT response")
	}

	// Debug istersen:
	// pretty, _ := json.MarshalIndent(raw, "", "  ")
	// fmt.Println(string(pretty))

	markets, ok := raw["markets"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("markets field missing in CCXT response")
	}

	pairs := make([]db.Pair, 0, len(markets))

	for symbol, entry := range markets {
		market, ok := entry.(map[string]interface{})
		if !ok {
			continue
		}

		base, _ := market["base"].(string)
		quote, _ := market["quote"].(string)

		// exchange’e tam uyumlu db.Pair struct
		pair := db.Pair{
			ExchangeID: exchange.ID,
			Symbol:     symbol,
			Base:       base,
			Quote:      quote,
		}

		pairs = append(pairs, pair)
	}

	log.Printf("[CEX] %s → %d pairs", exchange.Name, len(pairs))

	return pairs, nil
}
