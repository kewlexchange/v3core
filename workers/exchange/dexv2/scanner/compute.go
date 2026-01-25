package scanner

import (
	"core/models"
	"fmt"
	"math/big"
)

// Fee sabitleri
var (
	feeMul = big.NewInt(997)
	feeDiv = big.NewInt(1000)
)

func OptimalBorrow(
	RbIn, RbOut, RsIn, RsOut *big.Int,
) *big.Int {
	t := new(big.Int).Mul(RbIn, RbOut)
	t.Mul(t, RsIn)
	t.Mul(t, RsOut)
	t.Mul(t, feeMul)
	t.Mul(t, feeDiv)

	sqrtT := new(big.Int).Sqrt(t)

	sub := new(big.Int).Mul(RbIn, RsOut)
	sub.Mul(sub, feeDiv)

	if sqrtT.Cmp(sub) <= 0 {
		return big.NewInt(0)
	}

	den := new(big.Int).Mul(RsOut, feeMul)
	num := new(big.Int).Sub(sqrtT, sub)

	return num.Div(num, den)
}

func Price(pair models.Pair) float64 {
	q := new(big.Float).SetInt(pair.QuoteReserve)
	b := new(big.Float).SetInt(pair.BaseReserve)
	p, _ := new(big.Float).Quo(q, b).Float64()
	return p
}

// Rezervleri CHZ bazlı düzenle
func getReservesForCHZ(p models.Pair, chzAddress string, isBuy bool) (*big.Int, *big.Int) {
	if p.Base == chzAddress {
		if isBuy {
			// Buy pool'da CHZ base, input CHZ → RbIn=CHZ reserve, RbOut=diğer token reserve
			return p.BaseReserve, p.QuoteReserve
		} else {
			// Sell pool'da CHZ base, RsIn=diğer token reserve, RsOut=CHZ reserve
			return p.QuoteReserve, p.BaseReserve
		}
	} else if p.Quote == chzAddress {
		if isBuy {
			// Buy pool'da CHZ quote, input CHZ → RbIn=CHZ reserve, RbOut=diğer token reserve
			return p.QuoteReserve, p.BaseReserve
		} else {
			// Sell pool'da CHZ quote, RsIn=diğer token reserve, RsOut=CHZ reserve
			return p.BaseReserve, p.QuoteReserve
		}
	}
	// CHZ yoksa hata verebilirsin veya sıfır dönebilirsin
	return big.NewInt(0), big.NewInt(0)
}

func ComputeOptimalBorrowWithCHZOutput(scan models.ScanParams, pairs []models.Pair) (borrow *big.Int, buyPair models.Pair, sellPair models.Pair, err error) {
	if len(pairs) < 2 {
		return big.NewInt(0), models.Pair{}, models.Pair{}, fmt.Errorf("en az 2 pair gerekli")
	}

	minPrice := 1e50
	maxPrice := 0.0
	buyIndex, sellIndex := -1, -1

	for i, p := range pairs {
		// CHZ base mi quote mu ona göre fiyatı hesapla
		isChzBase := (p.Base == scan.NativeToken.Hex())
		var price float64

		if isChzBase {
			// CHZ base ise fiyat = Quote / Base
			price = Price(p)
		} else {
			// CHZ quote ise fiyat = Base / Quote
			// Ters fiyat hesapla
			price = 1.0 / Price(p)
		}

		if price < minPrice {
			minPrice = price
			buyIndex = i
		}
		if price > maxPrice {
			maxPrice = price
			sellIndex = i
		}
	}

	if buyIndex == sellIndex {
		return big.NewInt(0), models.Pair{}, models.Pair{}, fmt.Errorf("alım ve satım pool'ları aynı")
	}

	buy := pairs[buyIndex]
	sell := pairs[sellIndex]

	// CHZ'nin hangi tarafta olduğuna göre rezervleri ayarla
	buyRbIn, buyRbOut := getReservesForCHZ(buy, scan.NativeToken.Hex(), true)
	sellRsIn, sellRsOut := getReservesForCHZ(sell, scan.NativeToken.Hex(), false)

	borrow = OptimalBorrow(buyRbIn, buyRbOut, sellRsIn, sellRsOut)
	return borrow, buy, sell, nil
}

func CalculateProfit(borrow *big.Int, buy models.Pair, sell models.Pair, chzAddress string) *big.Int {
	buyRbIn, buyRbOut := getReservesForCHZ(buy, chzAddress, true)
	sellRsIn, sellRsOut := getReservesForCHZ(sell, chzAddress, false)

	buyAmountOut := getAmountOut(borrow, buyRbIn, buyRbOut)
	sellAmountOut := getAmountOut(buyAmountOut, sellRsIn, sellRsOut)

	return new(big.Int).Sub(sellAmountOut, borrow)
}

// getAmountOut = Uniswap formülüne göre çıktı hesaplama
func getAmountOut(amountIn, reserveIn, reserveOut *big.Int) *big.Int {
	amountInWithFee := new(big.Int).Mul(amountIn, big.NewInt(997))
	numerator := new(big.Int).Mul(amountInWithFee, reserveOut)
	denominator := new(big.Int).Add(new(big.Int).Mul(reserveIn, big.NewInt(1000)), amountInWithFee)
	return numerator.Div(numerator, denominator)
}

func Arbitrage(scan models.ScanParams, pairs []models.Pair) {

	borrow, buyPair, sellPair, err := ComputeOptimalBorrowWithCHZOutput(scan, pairs)

	if err != nil {
		return
	}

	// Borrow sınırla
	maxBorrow := new(big.Int).Div(buyPair.BaseReserve, big.NewInt(20)) // %5 sınır
	if borrow.Cmp(maxBorrow) > 0 {
		borrow = maxBorrow
	}

	profit := CalculateProfit(borrow, buyPair, sellPair, scan.NativeToken.Hex())

	msg := fmt.Sprintf("FROM %s -> TO %s, Borrow: %s Profit: %s", buyPair.Pair, sellPair.Pair, borrow.String(), profit.String())
	fmt.Println(msg)

	fmt.Printf("Buy Pair Reserves: Base=%s, Quote=%s\n", buyPair.BaseReserve.String(), buyPair.QuoteReserve.String())
	fmt.Printf("Sell Pair Reserves: Base=%s, Quote=%s\n", sellPair.BaseReserve.String(), sellPair.QuoteReserve.String())
	fmt.Printf("Borrow: %s\n", borrow.String())

	buyAmountOut := getAmountOut(borrow, buyPair.QuoteReserve, buyPair.BaseReserve)
	sellAmountOut := getAmountOut(buyAmountOut, sellPair.BaseReserve, sellPair.QuoteReserve)
	fmt.Printf("Buy Amount Out: %s\n", buyAmountOut.String())
	fmt.Printf("Sell Amount Out: %s\n", sellAmountOut.String())

	profitDebug := new(big.Int).Sub(sellAmountOut, borrow)
	fmt.Printf("Profit: %s\n", profitDebug.String())

}
