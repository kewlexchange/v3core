package exchange

import (
	"core/models"
	"fmt"
	"math"
	"strconv"
)

// Edge tipi Bellman-Ford için
type Edge struct {
	from, to int
	weight   float64
}

// Token listesini çıkarır ve index map döner
func getTokensFromPairs(pairs []models.Pair) ([]string, map[string]int) {
	tokenIndex := make(map[string]int)
	tokens := []string{}

	for _, p := range pairs {
		if _, ok := tokenIndex[p.Base]; !ok {
			tokenIndex[p.Base] = len(tokens)
			tokens = append(tokens, p.Base)
		}
		if _, ok := tokenIndex[p.Quote]; !ok {
			tokenIndex[p.Quote] = len(tokens)
			tokens = append(tokens, p.Quote)
		}
	}
	return tokens, tokenIndex
}

// Edge listesi oluşturur
func createEdgesFromPairs(pairs []models.Pair, tokenIndex map[string]int) []Edge {
	var edges []Edge

	for _, p := range pairs {
		if p.BasePrice == nil {
			continue
		}

		priceStr := p.BasePrice.String() // decimal.Decimal String() fonksiyonu var
		priceFloat, err := strconv.ParseFloat(priceStr, 64)
		if err != nil || priceFloat <= 0 {
			continue
		}

		from := tokenIndex[p.Base]
		to := tokenIndex[p.Quote]

		weight := -math.Log(priceFloat)
		edges = append(edges, Edge{from, to, weight})

		// Ters yön edge
		revWeight := -math.Log(1 / priceFloat)
		edges = append(edges, Edge{to, from, revWeight})
	}

	return edges
}

// Bellman-Ford ile negatif döngü bul ve cycle döndür
func findArbitrageCycle(edges []Edge, n int) ([]int, bool) {
	dist := make([]float64, n)
	predecessor := make([]int, n)
	for i := range predecessor {
		predecessor[i] = -1
	}

	for i := range dist {
		dist[i] = 0
	}

	var cycleStart int = -1

	for i := 0; i < n; i++ {
		updated := false
		for _, e := range edges {
			if dist[e.from]+e.weight < dist[e.to] {
				dist[e.to] = dist[e.from] + e.weight
				predecessor[e.to] = e.from
				updated = true
				if i == n-1 {
					cycleStart = e.to
				}
			}
		}
		if !updated {
			break
		}
	}

	if cycleStart == -1 {
		return nil, false
	}

	// Döngü başlangıcını n kez önceye taşı
	for i := 0; i < n; i++ {
		cycleStart = predecessor[cycleStart]
	}

	// Döngüyü çıkar
	cycle := []int{}
	current := cycleStart
	for {
		cycle = append(cycle, current)
		current = predecessor[current]
		if current == cycleStart && len(cycle) > 1 {
			break
		}
	}

	// Rotayı ters çevir
	for i, j := 0, len(cycle)-1; i < j; i, j = i+1, j-1 {
		cycle[i], cycle[j] = cycle[j], cycle[i]
	}

	return cycle, true
}

// Arbitraj tespiti ana fonksiyon
func DetectArbitrageOnPairs(pairs []models.Pair) {
	tokens, tokenIndex := getTokensFromPairs(pairs)
	edges := createEdgesFromPairs(pairs, tokenIndex)

	cycle, found := findArbitrageCycle(edges, len(tokens))
	if !found {
		fmt.Println("Arbitraj döngüsü bulunamadı.")
		return
	}

	fmt.Println("Arbitraj döngüsü bulundu:")
	for _, idx := range cycle {
		fmt.Println(tokens[idx])
	}
}
