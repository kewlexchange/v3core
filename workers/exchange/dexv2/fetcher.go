package dexv2

import (
	"core/models/db"
	"log"
	"time"
)

type DexV2Fetcher struct{}

func (d *DexV2Fetcher) FetchPairs(ex db.Exchange) ([]db.Pair, error) {
	// simulated work for DEX (on-chain calls, subgraph, etc.)
	log.Printf("[DEX Fetcher] fetching pairs for %s", ex.Name)
	time.Sleep(300 * time.Millisecond)
	pairs := []db.Pair{
		{Base: "TOKENA", Quote: "WETH"},
		{Base: "TOKENB", Quote: "WETH"},
	}
	log.Printf("[DEX Fetcher] fetched %d pairs for %s", len(pairs), ex.Name)
	return pairs, nil
}
