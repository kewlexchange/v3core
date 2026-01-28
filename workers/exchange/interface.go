package exchange

import (
	"core/models"
	coreTypes "core/types"
)

type PairFetcher interface {
	FetchPairs(ex models.Exchange) ([]models.TradingPair, error)
	FetchReserves(chainId int64, pairs []models.TradingPair) ([]models.TradingPair, error)
	ExecuteSwap(chainId int64, params coreTypes.ArbResult) error
}
