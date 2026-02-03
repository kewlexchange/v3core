package blockchains

import (
	"context"
	"core/constants"
	"core/models"
	"core/services"
	"core/workers"
	dexWorkers "core/workers/exchange/dexv2"
	"core/workers/exchange/dexv2/scanner/bsc"
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

type BSCScanner struct {
	client    *ethclient.Client
	rpcClient *rpc.Client

	scanning   atomic.Bool
	wg         sync.WaitGroup
	workerPool *workers.WorkerPool
	dexFetcher dexWorkers.DexV2Fetcher
	dexService *services.PairService
}

func NewBSCScanner(ctx context.Context, wsURL string) (*BSCScanner, error) {
	rpcClient, err := rpc.DialContext(ctx, wsURL)
	if err != nil {
		return nil, err
	}

	client := ethclient.NewClient(rpcClient)
	workerPool := workers.NewWorkerPool(50)
	dexFetcher := dexWorkers.DexV2Fetcher{}
	dexService := services.NewPairService(workerPool, &dexFetcher)

	return &BSCScanner{
		client:     client,
		rpcClient:  rpcClient,
		workerPool: workerPool,
		dexFetcher: dexFetcher,
		dexService: dexService,
	}, nil
}

func (s *BSCScanner) HandleScan(ctx context.Context) error {
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

	log.Printf("[HandleScanBSC] Starting dex scan with %d params\n", len(scanParamsBSC))

	s.dexService.ScanPairsSwapAll(constants.BSCChainId, scanParamsBSC)
	return nil
}

func (s *BSCScanner) Start(ctx context.Context) error {

	txs := make(chan *types.Transaction)

	sub, err := s.rpcClient.EthSubscribe(
		ctx,
		txs,
		"newPendingTransactions",
		true,
	)

	if err != nil {
		return err
	}

	log.Println("[BSCScanner] Started Pending Transaction subscription")

	for {
		select {

		case <-ctx.Done():
			log.Println("[BSCScanner] Context canceled")
			sub.Unsubscribe()
			return nil

		case err := <-sub.Err():
			return err

		case tx := <-txs:

			if tx.To() == nil {
				continue
			}

			if IsSwapMethod(tx) {

				log.Println("BSCScanner[BSCScanner]", tx.Hash().Hex())

				if s.scanning.CompareAndSwap(false, true) {
					go func() {
						defer s.scanning.Store(false)
						if err := s.HandleScan(ctx); err != nil {
							log.Printf("[BSCScanner] HandleScan error: %v", err)
						}
					}()
				}
			}
		}
	}
}

func (s *BSCScanner) Starz(ctx context.Context) error {
	parsedAbi, err := abi.JSON(strings.NewReader(constants.WETH_ABI))
	if err != nil {
		return err
	}

	contractAddress := constants.WETH_MAP[constants.BSCChainId][0]
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

	log.Println("[BSCScanner] Started WBNB Transfer event subscription")

	for {
		select {
		case <-ctx.Done():
			log.Println("[BSCScanner] Context canceled, unsubscribing")
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
				log.Printf("[BSCScanner] Failed to unpack transfer event: %v", err)
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
						log.Printf("[BSCScanner] HandleScan error: %v", err)
					}
				}()
			}
		}
	}
}

func (s *BSCScanner) Stop() {
	s.client.Close()
}
