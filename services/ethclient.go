package services

import (
	"core/models"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	clientsMu sync.Mutex
	Clients   = make(map[models.ChainID]*ethclient.Client)
)

func GetEVMClient(chainID models.ChainID, rpcURL string) (*ethclient.Client, error) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	if client, exists := Clients[chainID]; exists {
		return client, nil
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	Clients[chainID] = client
	return client, nil
}
