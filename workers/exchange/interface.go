package exchange

import "core/models/db"

type PairFetcher interface {
	FetchPairs(ex db.Exchange) ([]db.Pair, error)
}
