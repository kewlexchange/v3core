package services

import (
	"core/models"
	"core/workers"
	exchange "core/workers/exchange"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
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

func (s *PairService) SaveJSONToFile(outputDir, prefix string, data interface{}) (string, error) {
	// JSON'a çevir (indentli)
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("json marshal error: %w", err)
	}

	// Klasör oluştur
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("mkdir error: %w", err)
	}

	// Dosya adı oluştur (ör: output/prefix-<timestamp>.json)
	filename := filepath.Join(outputDir, fmt.Sprintf("%s-%d.json", prefix, time.Now().Unix()))

	// Dosyaya yaz
	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		return "", fmt.Errorf("write file error: %w", err)
	}

	return filename, nil
}

func (s *PairService) FetchPairsConcurrent(exchanges []models.Exchange) {
	for _, ex := range exchanges {
		exCopy := ex
		s.pool.Submit(func() {
			pairs, err := s.fetcher.FetchPairs(exCopy)
			if err != nil {
				log.Printf("error fetching for %s: %v", exCopy.Name, err)
				return
			}

			s.SaveJSONToFile("output", exCopy.Name, pairs)

			// in a real app persist pairs into DB or cache
			log.Printf("service received %d pairs from %s", len(pairs), exCopy.Name)
		})
	}
}
