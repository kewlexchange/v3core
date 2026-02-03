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
	"github.com/ethereum/go-ethereum/rpc"
)

type ChilizScanner struct {
	client    *ethclient.Client
	rpcClient *rpc.Client

	scanning   atomic.Bool
	wg         sync.WaitGroup
	workerPool *workers.WorkerPool
	dexFetcher dexWorkers.DexV2Fetcher
	dexService *services.PairService
}

func NewChilizScanner(ctx context.Context, wsURL string) (*ChilizScanner, error) {
	rpcClient, err := rpc.DialContext(ctx, wsURL)
	if err != nil {
		return nil, err
	}

	client := ethclient.NewClient(rpcClient)

	workerPool := workers.NewWorkerPool(50)
	dexFetcher := dexWorkers.DexV2Fetcher{}
	dexService := services.NewPairService(workerPool, &dexFetcher)

	return &ChilizScanner{
		client:     client,
		rpcClient:  rpcClient,
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

func (s *ChilizScanner) StartX(ctx context.Context) error {
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

func (s *ChilizScanner) Start(ctx context.Context) error {

	txs := make(chan *types.Transaction)

	sub, err := s.rpcClient.EthSubscribe(
		ctx,
		txs,
		"newPendingTransactions",
		true,
	)

	if err != nil {
		log.Println("ERR", err)
		return err
	}

	log.Println("[ChilizScanner] Started Pending Transaction subscription")

	for {
		select {

		case <-ctx.Done():
			log.Println("[ChilizScanner] Context canceled")
			sub.Unsubscribe()
			return nil

		case err := <-sub.Err():
			return err

		case tx := <-txs:

			if tx.To() == nil {
				continue
			}

			//if tx.To().Hex() == constants.WETH_MAP[constants.ChilizChainId][0].Hex() {

			//}

			//if IsSwapMethod(tx) {

			log.Println("PendingTxHashSwap[ChilizScanner]", tx.Hash())

			if s.scanning.CompareAndSwap(false, true) {
				go func() {
					defer s.scanning.Store(false)
					if err := s.HandleScan(ctx); err != nil {
						log.Printf("[ChilizScanner] HandleScan error: %v", err)
					}
				}()
			}
			//}
		}
	}
}

func (s *ChilizScanner) Stop() {
	s.client.Close()
}
