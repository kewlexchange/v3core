package services

import (
	"core/models/db"
	"core/workers"
	exchange "core/workers/exchange"
	"log"
)

type PairService struct {
	pool    *workers.WorkerPool
	fetcher exchange.PairFetcher
}

func NewPairService(pool *workers.WorkerPool, fetcher exchange.PairFetcher) *PairService {
	return &PairService{
		pool:    pool,
		fetcher: fetcher,
	}
}

func (s *PairService) FetchPairsConcurrent(exchanges []db.Exchange) {
	for _, ex := range exchanges {
		exCopy := ex
		s.pool.Submit(func() {
			pairs, err := s.fetcher.FetchPairs(exCopy)
			if err != nil {
				log.Printf("error fetching for %s: %v", exCopy.Name, err)
				return
			}
			// in a real app persist pairs into DB or cache
			log.Printf("service received %d pairs from %s", len(pairs), exCopy.Name)
		})
	}
}
