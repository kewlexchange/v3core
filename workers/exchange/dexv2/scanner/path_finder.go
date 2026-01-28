package scanner

import (
	"core/models"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type RouteStep struct {
	Base  common.Address
	Quote common.Address
	Pair  common.Address
}

func FindRoute(tradingPairs []models.TradingPair, input common.Address, output common.Address, maxHops int) ([]RouteStep, bool) {
	inputLower := input.Hex()
	outputLower := output.Hex()

	graph := make(map[string][]models.TradingPair)
	for _, pair := range tradingPairs {
		base := pair.BaseCurrency.Contract.Hex()
		quote := pair.QuoteCurrency.Contract.Hex()

		graph[strings.ToLower(base)] = append(graph[strings.ToLower(base)], pair)
		graph[strings.ToLower(quote)] = append(graph[strings.ToLower(quote)], pair)
	}

	type QueueItem struct {
		CurrentToken  string
		Path          []RouteStep
		VisitedTokens map[string]bool
	}

	queue := []QueueItem{{
		CurrentToken:  strings.ToLower(inputLower),
		Path:          []RouteStep{},
		VisitedTokens: map[string]bool{strings.ToLower(inputLower): true},
	}}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if len(item.Path) > maxHops {
			continue
		}

		if item.CurrentToken == strings.ToLower(outputLower) && len(item.Path) > 0 {
			return item.Path, true
		}

		for _, pair := range graph[item.CurrentToken] {
			base := strings.ToLower(pair.BaseCurrency.Contract.Hex())
			quote := strings.ToLower(pair.QuoteCurrency.Contract.Hex())

			var nextToken string
			if base == item.CurrentToken {
				nextToken = quote
			} else if quote == item.CurrentToken {
				nextToken = base
			} else {
				continue
			}

			if item.VisitedTokens[nextToken] {
				continue
			}

			newPath := make([]RouteStep, len(item.Path), len(item.Path)+1)
			copy(newPath, item.Path)
			newPath = append(newPath, RouteStep{
				Base:  common.HexToAddress(base),
				Quote: common.HexToAddress(quote),
				Pair:  pair.Pair,
			})

			newVisited := make(map[string]bool)
			for k, v := range item.VisitedTokens {
				newVisited[k] = v
			}
			newVisited[nextToken] = true

			queue = append(queue, QueueItem{
				CurrentToken:  nextToken,
				Path:          newPath,
				VisitedTokens: newVisited,
			})
		}
	}

	return nil, false
}
