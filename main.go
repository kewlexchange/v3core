package main

import (
	"core/constants"
	"core/services"
	"core/workers"
	"fmt"

	dexWorkers "core/workers/exchange/dexv2"
)

func main() {

	pool := workers.NewWorkerPool(100)

	// DEX fetcher
	dexFetcher := &dexWorkers.DexV2Fetcher{}
	dexService := services.NewPairService(pool, dexFetcher)

	// DEX exchanges

	for _, dex := range constants.DEXExchanges {
		client, err := services.GetEVMClient(*dex.ChainID, *dex.RPC)
		if err != nil {
			fmt.Println("GetEVMClient error:", err)
		}
		defer client.Close()
	}
	dexService.FetchPairsConcurrent(constants.DEXExchanges)

	/*
		// CEX exchanges
		cexExchanges := []models.Exchange{
			{Name: "BINANCE", Kind: models.ExchangeKindCEX},
			{Name: "BTCTURK", Kind: models.ExchangeKindCEX},
			{Name: "OKX", Kind: models.ExchangeKindCEX},
			{Name: "MEXC", Kind: models.ExchangeKindCEX},
		}
		for _, ex := range cexExchanges {

			switch ex.Name {

			case "BINANCE":

				client := ccxt.NewBinance(map[string]interface{}{
					"enableRateLimit": true,
				})
				fetcher := cexWorkers.NewCexFetcher(client) // POINTER gerekmez
				service := services.NewPairService(pool, fetcher)
				service.FetchPairsConcurrent([]models.Exchange{ex})

			case "MEXC":

				client := ccxt.NewMexc(map[string]interface{}{
					"enableRateLimit": true,
				})
				fetcher := cexWorkers.NewCexFetcher(client) // POINTER gerekmez
				service := services.NewPairService(pool, fetcher)
				service.FetchPairsConcurrent([]models.Exchange{ex})

			case "OKX":

				client := ccxt.NewOkx(map[string]interface{}{
					"enableRateLimit": true,
				})
				fetcher := cexWorkers.NewCexFetcher(client) // POINTER gerekmez
				service := services.NewPairService(pool, fetcher)
				service.FetchPairsConcurrent([]models.Exchange{ex})
			case "BTCTURK":
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
	*/

	pool.Wait()
}
