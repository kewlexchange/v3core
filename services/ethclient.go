package services

import (
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	clientsMu sync.Mutex
	clients   = make(map[int64]*ethclient.Client)
)

func GetEVMClient(chainID int64, rpcURL string) (*ethclient.Client, error) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	if client, exists := clients[chainID]; exists {
		return client, nil
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	clients[chainID] = client
	return client, nil
}
