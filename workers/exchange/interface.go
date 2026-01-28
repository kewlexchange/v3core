package exchange

import (
	"core/models"
	coreTypes "core/types"
)

type PairFetcher interface {
	FetchPairs(ex models.Exchange) ([]models.TradingPair, error)
	FetchReserves(chainId models.ChainID, pairs []models.TradingPair) ([]models.TradingPair, error)
	ExecuteSwap(chainId models.ChainID, params coreTypes.ArbResult) error
	ExecuteSwapAll(chainId models.ChainID, params []coreTypes.ArbResult) error
}
