package exchange

import "core/models"

type PairFetcher interface {
	FetchPairs(ex models.Exchange) ([]models.Pair, error)
}
