package exchange

import "core/models"

type PairFetcher interface {
	FetchPairs(ex models.Exchange) ([]models.TradingPair, error)
	FetchReserves(pairs []models.TradingPair) ([]models.TradingPair, error)
}
