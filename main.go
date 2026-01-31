package main

import (
	"context"
	"core/constants"
	"core/models"
	"core/services"
	blockchains "core/services/blockchains"
	"core/workers"
	dexWorkers "core/workers/exchange/dexv2"
	dexScanner "core/workers/exchange/dexv2/scanner"

	"core/workers/exchange/dexv2/scanner/avax"
	"core/workers/exchange/dexv2/scanner/bsc"
	"core/workers/exchange/dexv2/scanner/chiliz"
	"fmt"

	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
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
		avax.GetSWAP(),
		avax.GetDAI(),
		avax.GetCNR(),
		avax.GetSHRAP(),
		avax.GetPEFI(),
		avax.GetJEWEL(),
		avax.GetCRA(),
		avax.GetALOT(),
		avax.GetAVXT(),
		avax.GetYAK(),
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
		bsc.GetALPHA(),
		bsc.GetAXS(),
		bsc.GetALU(),
		bsc.GetBABYDOGE(),
		bsc.GetBMON(),
		bsc.GetBNX(),
		bsc.GetBSCPAD(),
		bsc.GetBTTOLD(),
		bsc.GetDPET(),
		bsc.GetMBOX(),
		bsc.GetINJ(),
		bsc.GetONE(),
		bsc.GetRACA(),
		bsc.GetREEF(),
	}

	fmt.Println("Len", len(scanParamsAVAX), len(scanParamsCHZ), len(scanParamsBSC))

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
			go dexService.ScanPairsSwapAll(constants.ChilizChainId, scanParamsCHZ)

			dexService.ScanPairsSwapAll(constants.AVAXChainID, scanParamsAVAX)
			//dexService.ScanPairsSwapAll(constants.BSCChainId, scanParamsBSC)
		}()
	}
}

func handleFetchCycles() {

	pool := workers.NewWorkerPool(50)

	dexFetcher := &dexWorkers.DexV2Fetcher{}
	dexService := services.NewPairService(pool, dexFetcher)
	cycles, err := dexService.FetchCycles(constants.ChilizChainId, chiliz.GetChilizCycle())

	if err != nil {
		fmt.Println("CYCLES:FAILED")
		return
	}

	for _, cycle := range cycles {
		dexScanner.TestCycle(constants.ChilizChainId, cycle)
	}

}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//handleFetchCycles()
	//return

	mainCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Chiliz için ayrı context
	chilizCtx, chilizCancel := context.WithCancel(mainCtx)
	defer chilizCancel()

	chilizScanner, err := blockchains.NewChilizScanner(chilizCtx, "wss://chiliz.publicnode.com")
	if err != nil {
		log.Printf("Failed to create Chiliz scanner: %v", err)
	}

	// Avalanche için ayrı context
	avaxCtx, avaxCancel := context.WithCancel(mainCtx)
	defer avaxCancel()

	avalancheScanner, err := blockchains.NewAvalancheScanner(avaxCtx, "wss://avalanche-c-chain-rpc.publicnode.com")
	if err != nil {
		fmt.Printf("Failed to create Avalanche scanner: %v", err)
	}

	bscCtx, bscCancel := context.WithCancel(mainCtx)
	defer bscCancel()
	bscScanner, err := blockchains.NewBSCScanner(bscCtx, "wss://bsc-rpc.publicnode.com")
	if err != nil {
		log.Fatalf("Failed to create BSC scanner: %v", err)
	}

	go func() {
		if err := chilizScanner.Start(chilizCtx); err != nil {
			log.Printf("Chiliz scanner error: %v", err)
		}
	}()

	go func() {
		return
		if err := bscScanner.Start(bscCtx); err != nil {
			log.Printf("bscScanner scanner error: %v", err)
		}
	}()

	go func() {
		if err := avalancheScanner.Start(avaxCtx); err != nil {
			log.Printf("Avalanche scanner error: %v", err)
		}
	}()
	/*


		go func() {
			if err := bscScanner.Start(bscCtx); err != nil {
				log.Printf("bscScanner scanner error: %v", err)
			}
		}()
	*/

	// OS sinyalleri yakala (Ctrl+C vs) graceful shutdown için
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	log.Println("Shutting down...")

	chilizScanner.Stop()
	avalancheScanner.Stop()
	bscScanner.Stop()

	// Biraz bekle tüm goroutine’lerin kapanması için
	time.Sleep(time.Second * 2)

	//workers.Start()

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
