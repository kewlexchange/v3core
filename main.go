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

	// DEX fetcher
	dexFetcher := &dexWorkers.DexV2Fetcher{}
	dexService := services.NewPairService(pool, dexFetcher)

	_ = dexService // kullanıyorsan

	// CEX exchanges
	cexExchanges := []db.Exchange{
		{Name: "Binance", Kind: db.ExchangeKindCEX},
		{Name: "BtcTurk", Kind: db.ExchangeKindCEX},
		{Name: "Paribu", Kind: db.ExchangeKindCEX},
	}

	for _, ex := range cexExchanges {

		switch ex.Name {

		case "Binance":

			client := ccxt.NewBinance(map[string]interface{}{
				"enableRateLimit": true,
			})
			fetcher := cexWorkers.NewCexFetcher(client) // POINTER gerekmez
			service := services.NewPairService(pool, fetcher)
			service.FetchPairsConcurrent([]db.Exchange{ex})

		case "BtcTurk":
			client := ccxt.NewBtcturk(map[string]interface{}{
				"enableRateLimit": true,
			})

			fetcher := cexWorkers.NewCexFetcher(client)
			service := services.NewPairService(pool, fetcher)
			service.FetchPairsConcurrent([]db.Exchange{ex})

		case "Paribu":
			println("[WARN] Paribu CCXT desteklemiyor, skip ediliyor.")
			// TODO: ParibuFetcher ekle
		}
	}
	pool.Wait()
}
