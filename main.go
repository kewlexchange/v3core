package main

import (
	"core/models/db"
	"core/services"
	"core/workers"

	cexWorkers "core/workers/exchange/cex"
	dexWorkers "core/workers/exchange/dexv2"

	ccxt "github.com/ccxt/ccxt/go/v4"
)

func main() {

	pool := workers.NewWorkerPool(100)
	dexFetcher := &dexWorkers.DexV2Fetcher{}
	cexFetcher := &cexWorkers.CexFetcher{}

	dexService := services.NewPairService(pool, dexFetcher)
	cexService := services.NewPairService(pool, cexFetcher)

	cexExchanges := []db.Exchange{
		{Name: "Binance", Kind: db.ExchangeKindCEX},
		{Name: "BtcTurk", Kind: db.ExchangeKindCEX},
		{Name: "Paribu", Kind: db.ExchangeKindCEX},
	}

	for _, ex := range cexExchanges {
		if ex.Name == "BtcTurk" {
			client := ccxt.NewBtcturk()
			service := cexService.NewCexFetcher(client)
			service.FetchPairsConcurrent([]db.Exchange{ex})
		}
		if ex.Name == "Binance" {
			client := ccxt.NewBtcturk()
			service := cex.NewCexFetcher(client)
			service.FetchPairsConcurrent([]db.Exchange{ex})
		}

	}

}
