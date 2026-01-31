package blockchains

import (
	"context"
	"core/constants"
	"core/models"
	"core/services"
	"core/workers"
	dexWorkers "core/workers/exchange/dexv2"
	"core/workers/exchange/dexv2/scanner/avax"
	"log"
	"math/big"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AvalancheScanner struct {
	client     *ethclient.Client
	scanning   atomic.Bool
	wg         sync.WaitGroup
	workerPool *workers.WorkerPool
	dexFetcher dexWorkers.DexV2Fetcher
	dexService *services.PairService
}

func NewAvalancheScanner(ctx context.Context, wsURL string) (*AvalancheScanner, error) {
	client, err := ethclient.DialContext(ctx, wsURL)
	if err != nil {
		return nil, err
	}

	workerPool := workers.NewWorkerPool(50)
	dexFetcher := dexWorkers.DexV2Fetcher{}
	dexService := services.NewPairService(workerPool, &dexFetcher)

	return &AvalancheScanner{
		client:     client,
		workerPool: workerPool,
		dexFetcher: dexFetcher,
		dexService: dexService,
	}, nil
}

func (s *AvalancheScanner) HandleScan(ctx context.Context) error {
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

	log.Printf("[HandleScan] Starting dex scan with %d params\n", len(scanParamsAVAX))

	s.dexService.ScanPairsSwapAll(constants.AVAXChainID, scanParamsAVAX)
	return nil
}

func (s *AvalancheScanner) Start(ctx context.Context) error {
	parsedAbi, err := abi.JSON(strings.NewReader(constants.WETH_ABI))
	if err != nil {
		return err
	}

	contractAddress := constants.WETH_MAP[constants.AVAXChainID][0]
	transferSigHash := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{transferSigHash}},
	}

	logs := make(chan types.Log)
	sub, err := s.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return err
	}

	log.Println("[AvalancheScanner] Started WAVAX Transfer event subscription")

	for {
		select {
		case <-ctx.Done():
			log.Println("[AvalancheScanner] Context canceled, unsubscribing")
			sub.Unsubscribe()
			return nil
		case err := <-sub.Err():
			return err
		case vLog := <-logs:
			var transferEvent struct {
				From  common.Address
				To    common.Address
				Value *big.Int
			}

			err := parsedAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Printf("[AvalancheScanner] Failed to unpack transfer event: %v", err)
				continue
			}

			// indexed alanlar topics'ten
			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			/*
				log.Printf("[BSCScanner] WBNB Transfer detected: From %s To %s Value %s TxHash %s",
					transferEvent.From.Hex(),
					transferEvent.To.Hex(),
					transferEvent.Value.String(),

			*/

			if s.scanning.CompareAndSwap(false, true) {
				go func() {
					defer s.scanning.Store(false)
					if err := s.HandleScan(ctx); err != nil {
						log.Printf("[AvalancheScanner] HandleScan error: %v", err)
					}
				}()
			}
		}
	}
}

func (s *AvalancheScanner) StartEx(ctx context.Context) error {
	headers := make(chan *types.Header)
	sub, err := s.client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return err
	}
	log.Println("[AvalancheScanner] Started")

	for {
		select {
		case <-ctx.Done():
			log.Println("[AvalancheScanner] Context canceled")
			return nil
		case err := <-sub.Err():
			return err
		case header := <-headers:

			log.Printf("[AvalancheScanner] New block: %d\n", header.Number.Uint64())

			if s.scanning.CompareAndSwap(false, true) {
				go func() {
					defer s.scanning.Store(false)
					if err := s.HandleScan(ctx); err != nil {
						log.Printf("[HandleScan] AVALANCHE error: %v\n", err)
					}
				}()
			} else {
				log.Println("[HandleScan] AVALANCHE already running, skipping")
			}
		}
	}
}

func (s *AvalancheScanner) Stop() {
	s.client.Close()
}
