package blockchains

import (
	"context"
	"core/constants"
	"core/models"
	"core/services"
	"core/workers"
	dexWorkers "core/workers/exchange/dexv2"
	"core/workers/exchange/dexv2/scanner/chiliz"
	"fmt"
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

type ChilizScanner struct {
	client     *ethclient.Client
	scanning   atomic.Bool
	wg         sync.WaitGroup
	workerPool *workers.WorkerPool
	dexFetcher dexWorkers.DexV2Fetcher
	dexService *services.PairService
}

func NewChilizScanner(ctx context.Context, wsURL string) (*ChilizScanner, error) {
	client, err := ethclient.DialContext(ctx, wsURL)
	if err != nil {
		return nil, err
	}

	workerPool := workers.NewWorkerPool(50)
	dexFetcher := dexWorkers.DexV2Fetcher{}
	dexService := services.NewPairService(workerPool, &dexFetcher)

	return &ChilizScanner{
		client:     client,
		workerPool: workerPool,
		dexFetcher: dexFetcher,
		dexService: dexService,
	}, nil
}

func (s *ChilizScanner) HandleScan(ctx context.Context) error {
	scanParamsCHZ := []models.ScanParams{
		chiliz.GetUSDC(),
		chiliz.GetPEPPER(),
		chiliz.GetCHZINU(),
	}

	log.Printf("[HandleScanCHZ] Starting dex scan with %d params\n", len(scanParamsCHZ))

	s.dexService.ScanPairsSwapAll(constants.ChilizChainId, scanParamsCHZ)
	return nil
}

func (s *ChilizScanner) Start(ctx context.Context) error {
	parsedAbi, err := abi.JSON(strings.NewReader(constants.WETH_ABI))
	if err != nil {
		return err
	}

	contractAddress := constants.WETH_MAP[constants.ChilizChainId][0]
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

	log.Println("[ChilizScanner] Started WCHZ Transfer event subscription")

	for {
		select {
		case <-ctx.Done():
			log.Println("[ChilizScanner] Context canceled, unsubscribing")
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
				log.Printf("[ChilizScanner] Failed to unpack transfer event: %v", err)
				continue
			}

			// indexed alanlar topics'ten
			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Println("GelenTxHash", vLog.TxHash)

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
						log.Printf("[ChilizScanner] HandleScan error: %v", err)
					}
				}()
			}
		}
	}
}

func (s *ChilizScanner) StartEx(ctx context.Context) error {
	headers := make(chan *types.Header)
	sub, err := s.client.SubscribeNewHead(ctx, headers)
	if err != nil {
		return err
	}
	log.Println("[ChilizScanner] Started")

	for {
		select {
		case <-ctx.Done():
			log.Println("[ChilizScanner] Context canceled")
			return nil
		case err := <-sub.Err():
			return err
		case header := <-headers:
			log.Printf("[ChilizScanner] New block: %d\n", header.Number.Uint64())
			block, err := s.client.BlockByHash(ctx, header.Hash())
			if err != nil {
				log.Printf("[ChilizScanner] Failed to fetch block: %v\n", err)
				continue
			}

			var foundSwap atomic.Bool
			s.wg = sync.WaitGroup{}

			for _, tx := range block.Transactions() {
				if tx.To() == nil {
					continue
				}
				s.wg.Add(1)
				go func(tx *types.Transaction) {
					defer s.wg.Done()
					if IsSwapMethod(tx) {
						log.Printf("[ChilizScanner] Swap detected: %s\n", tx.Hash().Hex())
						foundSwap.Store(true)
					}
				}(tx)
			}

			s.wg.Wait()

			if foundSwap.Load() {
				if s.scanning.CompareAndSwap(false, true) {
					go func() {
						defer s.scanning.Store(false)
						if err := s.HandleScan(ctx); err != nil {
							log.Printf("[StartCHZ] HandleScanCHZ error: %v\n", err)
						}
					}()
				} else {
					log.Println("[StartCHZ] HandleScanCHZ already running, skipping")
				}
			}
		}
	}
}

func (s *ChilizScanner) Stop() {
	s.client.Close()
}
