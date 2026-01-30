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
	"log"
	"sync/atomic"
	"time"

	"github.com/joho/godotenv"
)

func handleScan() {
	pool := workers.NewWorkerPool(50)

	dexFetcher := &dexWorkers.DexV2Fetcher{}
	dexService := services.NewPairService(pool, dexFetcher)

	scanParamsAVAX := []models.ScanParams{
		avax.GetROCO(),
		avax.GetUSDC(),
		avax.GetJOE(),
		avax.GetXAVA(),
		avax.GetCOQ(),
		avax.GetPNG(),
		avax.GetARENA(),
		avax.GetUSDT(),
		avax.GetAAAVE(),
		avax.GetWETH(),
		avax.GetSAVAX(),
		avax.GetUSDCE(),
		avax.GetTSD(),
		avax.GetLINK(),
		avax.GetSUSHI(),
		avax.GetUSDTE(),
		avax.GetDAI(),
	}

	scanParamsCHZ := []models.ScanParams{
		chiliz.GetUSDC(),
		chiliz.GetPEPPER(),
		chiliz.GetCHZINU(),
	}

	scanParamsBSC := []models.ScanParams{
		bsc.GetUSDC(),
		bsc.GetUSDT(),
		bsc.GetCAKE(),
		bsc.GetETH(),
		bsc.GetBTCB(),
		bsc.GetADA(),
		bsc.GetAUTO(),
		bsc.GetBAT(),
		bsc.GetBSW(),
		bsc.GetC98(),
		bsc.GetCEEK(),
		bsc.GetDOGE(),
	}

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	var scanning atomic.Bool

	for range ticker.C {
		if scanning.Load() {
			continue
		}

		scanning.Store(true)
		go func() {
			defer scanning.Store(false)
			go dexService.ScanPairsSwapAll(constants.AVAXChainID, scanParamsAVAX)
			go dexService.ScanPairsSwapAll(constants.ChilizChainId, scanParamsCHZ)
			go dexService.ScanPairsSwapAll(constants.BSCChainId, scanParamsBSC)
		}()
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	handleScan()
	//services.FetchBalancesFromPKEY()

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

	   pool.Wait()
	*/
}
