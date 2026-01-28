package scanner

import (
	"core/models"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestFindRoute(t *testing.T) {
	chz := common.HexToAddress("0xf3928e7871eb136dd6648ad08aeef6b6ea893001")
	tokenA := common.HexToAddress("0x1111111111111111111111111111111111111111")
	tokenB := common.HexToAddress("0x2222222222222222222222222222222222222222")
	tokenC := common.HexToAddress("0x3333333333333333333333333333333333333333")
	tokenD := common.HexToAddress("0x4444444444444444444444444444444444444444")

	pairs := []models.TradingPair{
		{
			BaseCurrency:  models.Currency{Contract: &chz},
			QuoteCurrency: models.Currency{Contract: &tokenA},
			Pair:          common.HexToAddress("0xaaaa000000000000000000000000000000000000"),
		},
		{
			BaseCurrency:  models.Currency{Contract: &tokenA},
			QuoteCurrency: models.Currency{Contract: &tokenB},
			Pair:          common.HexToAddress("0xbbbb000000000000000000000000000000000000"),
		},
		{
			BaseCurrency:  models.Currency{Contract: &tokenB},
			QuoteCurrency: models.Currency{Contract: &tokenC},
			Pair:          common.HexToAddress("0xcccc000000000000000000000000000000000000"),
		},
		{
			BaseCurrency:  models.Currency{Contract: &tokenC},
			QuoteCurrency: models.Currency{Contract: &tokenD},
			Pair:          common.HexToAddress("0xdddd000000000000000000000000000000000000"),
		},
	}

	input := chz
	output := tokenD
	maxHops := 5

	route, found := FindRoute(pairs, input, output, maxHops)
	if !found {
		t.Fatalf("Route not found but expected to find one")
	}

	expectedLength := 4
	if len(route) != expectedLength {
		t.Fatalf("Expected route length %d, got %d", expectedLength, len(route))
	}

	// Check the sequence of pairs
	expectedPairs := []common.Address{
		common.HexToAddress("0xaaaa000000000000000000000000000000000000"),
		common.HexToAddress("0xbbbb000000000000000000000000000000000000"),
		common.HexToAddress("0xcccc000000000000000000000000000000000000"),
		common.HexToAddress("0xdddd000000000000000000000000000000000000"),
	}

	for i, step := range route {
		if step.Pair != expectedPairs[i] {
			t.Errorf("Step %d: expected pair %s, got %s", i, expectedPairs[i].Hex(), step.Pair.Hex())
		}
	}

	fmt.Println("Input", input.Hex(), "Output", output.Hex())
	for i, step := range route {
		t.Logf("Step %d: Pair=%s, Base=%s, Quote=%s", i, step.Pair.Hex(), step.Base.Hex(), step.Quote.Hex())
	}
}
