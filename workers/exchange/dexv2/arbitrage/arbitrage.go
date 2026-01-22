package arbitrage

import (
	"core/models"
	"fmt"
	"math/big"
)

// BI converts a big.Int to a big.Float.
func BI(x *big.Int) *big.Float {
	return new(big.Float).SetInt(x)
}

var (
	// Gamma represents the trading fee (0.3%), so the multiplier is 1 - 0.003 = 0.997.
	Gamma = big.NewFloat(0.997)
)

// SpotPrice calculates the spot price of a pair (Quote/Base).
func SpotPrice(p *models.Pair) *big.Float {
	if p.BaseReserve.Sign() == 0 {
		return big.NewFloat(0)
	}
	return new(big.Float).Quo(
		BI(p.QuoteReserve),
		BI(p.BaseReserve),
	)
}

// PickBuySell identifies the pools with the lowest and highest prices.
func PickBuySell(pairs []*models.Pair) (buy, sell *models.Pair) {
	if len(pairs) == 0 {
		return nil, nil
	}

	buy = pairs[0]
	sell = pairs[0]

	for _, p := range pairs {
		if SpotPrice(p).Cmp(SpotPrice(buy)) < 0 {
			buy = p
		}
		if SpotPrice(p).Cmp(SpotPrice(sell)) > 0 {
			sell = p
		}
	}

	if buy == sell {
		return nil, nil
	}

	return buy, sell
}

// --- Arbitrage Cycle: quote -> base -> quote ---

// OptimalBorrowQuote calculates the optimal amount of quote token for a "borrow quote" cycle.
// This cycle is profitable if buyPrice < sellPrice.
// Formula source: derived from maximizing profit function for two-pool arbitrage.
func OptimalBorrowQuote(buy, sell *models.Pair) *big.Float {
	xb := BI(buy.BaseReserve)
	yb := BI(buy.QuoteReserve)
	xs := BI(sell.BaseReserve)
	ys := BI(sell.QuoteReserve)
	g := Gamma
	g2 := new(big.Float).Mul(g, g)

	// Profitability condition: (ys/xs)*g2 > (yb/xb)  =>  ys*xb*g2 > yb*xs
	cond1 := new(big.Float).Mul(ys, xb)
	cond1.Mul(cond1, g2)
	cond2 := new(big.Float).Mul(yb, xs)

	if cond1.Cmp(cond2) <= 0 {
		return big.NewFloat(0)
	}

	// Formula: (sqrt(xb*xs*yb*ys*g*g) - yb*xs) / (yb*g + ys)
	sqrtNumerator := new(big.Float).Mul(xb, xs)
	sqrtNumerator.Mul(sqrtNumerator, yb)
	sqrtNumerator.Mul(sqrtNumerator, ys)
	sqrtNumerator.Mul(sqrtNumerator, g2)
	sqrtNumerator.Sqrt(sqrtNumerator)

	numerator := new(big.Float).Sub(sqrtNumerator, cond2)

	// Denominator: yb*g + ys
	denominator := new(big.Float).Mul(yb, g)
	denominator.Add(denominator, ys)

	if denominator.Sign() == 0 {
		return big.NewFloat(0)
	}

	return new(big.Float).Quo(numerator, denominator)
}

// ProfitQuote calculates the profit in quote token from a "borrow quote" cycle.
// amountIn is the optimal amount of quote token.
func ProfitQuote(buy, sell *models.Pair, amountIn *big.Float) *big.Float {
	xb := BI(buy.BaseReserve)
	yb := BI(buy.QuoteReserve)
	xs := BI(sell.BaseReserve)
	ys := BI(sell.QuoteReserve)

	// 1. Buy base from 'buy' pool with quote token (amountIn)
	// bOut = (amountIn * gamma * xb) / (yb + amountIn * gamma)
	bOutNum := new(big.Float).Mul(new(big.Float).Mul(amountIn, Gamma), xb)
	bOutDen := new(big.Float).Add(yb, new(big.Float).Mul(amountIn, Gamma))
	bOut := new(big.Float).Quo(bOutNum, bOutDen)

	// 2. Sell base (bOut) to 'sell' pool for quote token
	// qBack = (bOut * gamma * ys) / (xs + bOut * gamma)
	qBackNum := new(big.Float).Mul(new(big.Float).Mul(bOut, Gamma), ys)
	qBackDen := new(big.Float).Add(xs, new(big.Float).Mul(bOut, Gamma))
	qBack := new(big.Float).Quo(qBackNum, qBackDen)

	// 3. Profit = qBack - amountIn
	return new(big.Float).Sub(qBack, amountIn)
}

// --- Arbitrage Cycle: base -> quote -> base ---

// OptimalBorrowBase calculates the optimal amount of base token for a "borrow base" cycle.
// This cycle is profitable if buyPrice > sellPrice.
func OptimalBorrowBase(buy, sell *models.Pair) *big.Float {
	xb := BI(buy.BaseReserve)
	yb := BI(buy.QuoteReserve)
	xs := BI(sell.BaseReserve)
	ys := BI(sell.QuoteReserve)
	g := Gamma
	g2 := new(big.Float).Mul(g, g)

	// Profitability condition: (yb/xb)*g2 > (ys/xs)  =>  yb*xs*g2 > xb*ys
	cond1 := new(big.Float).Mul(yb, xs)
	cond1.Mul(cond1, g2)
	cond2 := new(big.Float).Mul(xb, ys)

	if cond1.Cmp(cond2) <= 0 {
		return big.NewFloat(0)
	}

	// Formula: (sqrt(xb*xs*yb*ys*g*g) - xb*ys) / (xb*g + xs)
	sqrtNumerator := new(big.Float).Mul(xb, xs)
	sqrtNumerator.Mul(sqrtNumerator, yb)
	sqrtNumerator.Mul(sqrtNumerator, ys)
	sqrtNumerator.Mul(sqrtNumerator, g2)
	sqrtNumerator.Sqrt(sqrtNumerator)

	numerator := new(big.Float).Sub(sqrtNumerator, cond2)

	// Denominator: xb*g + xs
	denominator := new(big.Float).Mul(xb, g)
	denominator.Add(denominator, xs)

	if denominator.Sign() == 0 {
		return big.NewFloat(0)
	}

	return new(big.Float).Quo(numerator, denominator)
}

// ProfitBase calculates the profit in base token from a "borrow base" cycle.
// amountIn is the optimal amount of base token.
func ProfitBase(buy, sell *models.Pair, amountIn *big.Float) *big.Float {
	xb := BI(buy.BaseReserve)
	yb := BI(buy.QuoteReserve)
	xs := BI(sell.BaseReserve)
	ys := BI(sell.QuoteReserve)

	// 1. Sell base (amountIn) to 'sell' pool for quote token
	// qOut = (amountIn * gamma * ys) / (xs + amountIn * gamma)
	qOutNum := new(big.Float).Mul(new(big.Float).Mul(amountIn, Gamma), ys)
	qOutDen := new(big.Float).Add(xs, new(big.Float).Mul(amountIn, Gamma))
	qOut := new(big.Float).Quo(qOutNum, qOutDen)

	// 2. Buy base from 'buy' pool with quote token (qOut)
	// bBack = (qOut * gamma * xb) / (yb + qOut * gamma)
	bBackNum := new(big.Float).Mul(new(big.Float).Mul(qOut, Gamma), xb)
	bBackDen := new(big.Float).Add(yb, new(big.Float).Mul(qOut, Gamma))
	bBack := new(big.Float).Quo(bBackNum, bBackDen)

	// 3. Profit = bBack - amountIn
	return new(big.Float).Sub(bBack, amountIn)
}

func TestFunctions() {

	testCases := []struct {
		name  string
		pairs []*models.Pair
	}{
		{
			name: "Borrow Quote - Price difference large",
			pairs: []*models.Pair{
				{
					Pair:         "TOKENA/TOKENB",
					Base:         "TOKENA",
					Quote:        "TOKENB",
					BaseReserve:  big.NewInt(1_000_000_000_000_000_000), // 1 token
					QuoteReserve: big.NewInt(500_000_000_000_000_000),   // 0.5 price
				},
				{
					Pair:         "TOKENA/TOKENB",
					Base:         "TOKENA",
					Quote:        "TOKENB",
					BaseReserve:  big.NewInt(500_000_000_000_000_000),   // 0.5 token
					QuoteReserve: big.NewInt(2_000_000_000_000_000_000), // 4.0 price
				},
			},
		},
		{
			name: "Borrow Base - Reverse price difference",
			pairs: []*models.Pair{
				{
					Pair:         "TOKENA/TOKENB",
					Base:         "TOKENA",
					Quote:        "TOKENB",
					BaseReserve:  big.NewInt(500_000_000_000_000_000),   // 0.5 token
					QuoteReserve: big.NewInt(2_000_000_000_000_000_000), // price 4.0
				},
				{
					Pair:         "TOKENA/TOKENB",
					Base:         "TOKENA",
					Quote:        "TOKENB",
					BaseReserve:  big.NewInt(1_000_000_000_000_000_000), // 1 token
					QuoteReserve: big.NewInt(500_000_000_000_000_000),   // price 0.5
				},
			},
		},
		{
			name: "No Arbitrage - Equal prices",
			pairs: []*models.Pair{
				{
					Pair:         "TOKENA/TOKENB",
					Base:         "TOKENA",
					Quote:        "TOKENB",
					BaseReserve:  big.NewInt(1_000_000_000_000_000_000),
					QuoteReserve: big.NewInt(1_000_000_000_000_000_000),
				},
				{
					Pair:         "TOKENA/TOKENB",
					Base:         "TOKENA",
					Quote:        "TOKENB",
					BaseReserve:  big.NewInt(1_000_000_000_000_000_000),
					QuoteReserve: big.NewInt(1_000_000_000_000_000_000),
				},
			},
		},
	}

	for _, tc := range testCases {
		fmt.Println("====================================")
		fmt.Println("Test Case:", tc.name)
		FindArbitrage(tc.pairs)
		fmt.Println("====================================")
	}

	fmt.Println("END_TEST")
}
func FindArbitrage(pairs []*models.Pair) {

	buy, sell := PickBuySell(pairs)
	if buy == nil || sell == nil {
		fmt.Println("Arbitrage opportunity not found (not enough pairs or no price difference).")
		return
	}

	buyPrice := SpotPrice(buy)
	sellPrice := SpotPrice(sell)

	// Divisor to convert from wei-like units to whole token units for display
	divisor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

	// Case 1: Borrow Quote (quote -> base -> quote)
	// Profitable if we can buy cheap and sell expensive: buyPrice < sellPrice
	if buyPrice.Cmp(sellPrice) < 0 {
		fmt.Println("Checking for 'Borrow Quote' arbitrage...")
		fmt.Printf("Buy from: %s (Price: %s)\n", buy.Pair, buyPrice.Text('f', 18))
		fmt.Printf("Sell to: %s (Price: %s)\n", sell.Pair, sellPrice.Text('f', 18))

		borrowAmount := OptimalBorrowQuote(buy, sell)
		if borrowAmount.Sign() <= 0 {
			fmt.Println("Arbitrage not profitable.")
			return
		}

		profit := ProfitQuote(buy, sell, borrowAmount)
		if profit.Sign() <= 0 {
			fmt.Println("No net profit after calculation.")
			return
		}

		// --- Format output for readability ---
		displayBorrow := new(big.Float).Quo(borrowAmount, divisor)
		displayProfit := new(big.Float).Quo(profit, divisor)
		borrowInt, _ := borrowAmount.Int(nil)
		profitInt, _ := profit.Int(nil)

		fmt.Println("\n--- ARBITRAGE FOUND (Borrow Quote) ---")
		fmt.Printf("Buy Pool:  %s\n", buy.Pair)
		fmt.Printf("Sell Pool: %s\n", sell.Pair)
		fmt.Println("--")
		fmt.Println("NOTE: Amounts are displayed in whole tokens (divided by 10^18).")
		fmt.Printf("Borrow Amount: %s tokens\n", displayBorrow.Text('f', 18))
		fmt.Printf("Profit:        %s tokens\n", displayProfit.Text('f', 18))
		fmt.Println("--")
		fmt.Printf("Raw Borrow (units): %s\n", borrowInt.String())
		fmt.Printf("Raw Profit (units): %s\n", profitInt.String())
		return
	}

	// Case 2: Borrow Base (base -> quote -> base)
	// Profitable if buyPrice > sellPrice
	if buyPrice.Cmp(sellPrice) > 0 {
		fmt.Println("Checking for 'Borrow Base' arbitrage...")
		// For this cycle, we sell base at the 'buy' pool (where its price is high)
		// and buy it back at the 'sell' pool (where its price is low).
		fmt.Printf("Sell Base Token at: %s (High Base Price: %s)\n", buy.Pair, buyPrice.Text('f', 18))
		fmt.Printf("Buy Base Token at: %s (Low Base Price: %s)\n", sell.Pair, sellPrice.Text('f', 18))

		borrowAmount := OptimalBorrowBase(buy, sell)
		if borrowAmount.Sign() <= 0 {
			fmt.Println("Arbitrage not profitable.")
			return
		}

		// To calculate profit, we use the original buy/sell pools but correct function
		profit := ProfitBase(buy, sell, borrowAmount)
		if profit.Sign() <= 0 {
			fmt.Println("No net profit after calculation.")
			return
		}

		// --- Format output for readability ---
		displayBorrow := new(big.Float).Quo(borrowAmount, divisor)
		displayProfit := new(big.Float).Quo(profit, divisor)
		borrowInt, _ := borrowAmount.Int(nil)
		profitInt, _ := profit.Int(nil)

		fmt.Println("\n--- ARBITRAGE FOUND (Borrow Base) ---")
		fmt.Printf("High Price Pool (Sell Base): %s\n", buy.Pair)
		fmt.Printf("Low Price Pool (Buy Base):  %s\n", sell.Pair)
		fmt.Println("--")
		fmt.Println("NOTE: Amounts are displayed in whole tokens (divided by 10^18).")
		fmt.Printf("Borrow Amount: %s tokens\n", displayBorrow.Text('f', 18))
		fmt.Printf("Profit:        %s tokens\n", displayProfit.Text('f', 18))
		fmt.Println("--")
		fmt.Printf("Raw Borrow (units): %s\n", borrowInt.String())
		fmt.Printf("Raw Profit (units): %s\n", profitInt.String())
		return
	}

	fmt.Println("No profitable arbitrage opportunity found.")
}
