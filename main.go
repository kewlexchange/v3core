package main

import (
	"core/constants"
	"core/models"
	"core/services"
	"core/workers"
	dexWorkers "core/workers/exchange/dexv2"
	"core/workers/exchange/dexv2/scanner/avax"
	"core/workers/exchange/dexv2/scanner/bsc"
	"core/workers/exchange/dexv2/scanner/chiliz"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//services.FetchBalancesFromPKEY()

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

	scanParamsCHZ := []models.ScanParams{
		chiliz.GetUSDC(),
		chiliz.GetPEPPER(),
		chiliz.GetCHZINU(),
		//chiliz.GetComplex(),
	}

	scanParamsBSC := []models.ScanParams{
		bsc.GetUSDC(),
		bsc.GetUSDT(),
		bsc.GetCAKE(),
		//bsc.GetComplex(),
	}

	scanParamsAVAX := []models.ScanParams{
		avax.GetROCO(),
		avax.GetUSDC(),
	}

	dexService.ScanPairsSwapAll(constants.ChilizChainId, scanParamsCHZ)
	dexService.ScanPairsSwapAll(constants.BSCChainId, scanParamsBSC)
	dexService.ScanPairsSwapAll(constants.AVAXChainID, scanParamsAVAX)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C

		scanParamsCHZ := []models.ScanParams{
			chiliz.GetPEPPER(),
			chiliz.GetCHZINU(),
			chiliz.GetUSDC(),
			//chiliz.GetComplex(),
		}
		scanParamsBSC := []models.ScanParams{
			bsc.GetUSDC(),
			bsc.GetUSDT(),
			bsc.GetCAKE(),
			//bsc.GetComplex(),
		}

		scanParamsAVAX := []models.ScanParams{
			avax.GetROCO(),
			avax.GetUSDC(),
		}

		dexService.ScanPairSwapSingle(constants.ChilizChainId, scanParamsCHZ)
		dexService.ScanPairSwapSingle(constants.BSCChainId, scanParamsBSC)
		dexService.ScanPairSwapSingle(constants.AVAXChainID, scanParamsAVAX)
	}
	//dexService.FetchPairsConcurrent(constants.DEXExchanges)

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


	pool.Wait()*/
}
