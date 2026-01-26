package scanner

import (
	"core/models"
	"core/utils"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Fee sabitleri (Uniswap benzeri %0.3 komisyon)
var (
	feeMul = big.NewInt(997)
	feeDiv = big.NewInt(1000)
)

func ComputeBorrowAmount(RbIn, RbOut, RsIn, RsOut *big.Int) *big.Int {
	feeMul := big.NewInt(997)
	feeDiv := big.NewInt(1000)

	feeMulF := new(big.Float).SetInt(feeMul)
	feeDivF := new(big.Float).SetInt(feeDiv)

	RbInF := new(big.Float).SetInt(RbIn)
	RbOutF := new(big.Float).SetInt(RbOut)
	RsInF := new(big.Float).SetInt(RsIn)
	RsOutF := new(big.Float).SetInt(RsOut)

	// t = RbIn * RbOut * RsIn * RsOut * feeMul / feeDiv
	t := new(big.Float).Mul(RbInF, RbOutF)
	t.Mul(t, RsInF)
	t.Mul(t, RsOutF)
	t.Mul(t, feeMulF)
	t.Quo(t, feeDivF)

	sqrtT := new(big.Float).Sqrt(t)

	sub := new(big.Float).Mul(RbInF, RsOutF)

	diff := new(big.Float).Sub(sqrtT, sub)
	zero := big.NewFloat(0)
	if diff.Cmp(zero) <= 0 {
		return big.NewInt(0)
	}

	denom := new(big.Float).Mul(RsOutF, feeMulF)
	denom.Quo(denom, feeDivF)

	borrowF := new(big.Float).Quo(diff, denom)

	borrowInt, _ := borrowF.Int(nil)
	if borrowInt.Cmp(big.NewInt(0)) < 0 {
		return big.NewInt(0)
	}

	return borrowInt
}

func PriceCHZPerToken(p models.Pair, nativeToken common.Address) float64 {
	// decimal değerlerini int'e çevir

	var reserveIn, reserveOut *big.Int
	var decimals int64

	if p.Base == nativeToken.Hex() {
		reserveIn = p.BaseReserve
		reserveOut = p.QuoteReserve
		decimals = p.BaseDecimals.Int64() // Base token decimal
	} else if p.Quote == nativeToken.Hex() {
		reserveIn = p.QuoteReserve
		reserveOut = p.BaseReserve
		decimals = p.QuoteDecimals.Int64() // Quote token decimal
	} else {
		return 0
	}

	// one = 10^decimals
	one := new(big.Int).Exp(big.NewInt(10), big.NewInt(decimals), nil)

	amountOut := getAmountOut(one, reserveIn, reserveOut)
	if amountOut.Cmp(big.NewInt(0)) == 0 {
		return 0
	}

	amountOutF := new(big.Float).SetInt(amountOut)
	oneF := new(big.Float).SetInt(one)

	priceFloat, _ := new(big.Float).Quo(amountOutF, oneF).Float64()

	return priceFloat
}

// getReservesForCHZ, pair'daki rezervleri CHZ bazlı input-output olarak döner.
// isBuy: true ise input token CHZ, false ise output token CHZ
func getReservesForCHZ(p models.Pair, chzAddress string, isBuy bool) (*big.Int, *big.Int) {
	if isBuy {
		if p.Base == chzAddress {
			return p.BaseReserve, p.QuoteReserve
		} else if p.Quote == chzAddress {
			return p.QuoteReserve, p.BaseReserve
		}
	} else {
		if p.Base == chzAddress {
			return p.QuoteReserve, p.BaseReserve
		} else if p.Quote == chzAddress {
			return p.BaseReserve, p.QuoteReserve
		}
	}
	return big.NewInt(0), big.NewInt(0)
}

func ComputeOptimalBorrowWithCHZOutput(scan models.ScanParams, pairs []models.Pair) (*big.Int, models.Pair, models.Pair, error) {
	if len(pairs) < 2 {
		return big.NewInt(0), models.Pair{}, models.Pair{}, fmt.Errorf("en az 2 pair gerekli")
	}

	minPrice := 1e50
	maxPrice := 0.0
	buyIndex, sellIndex := -1, -1

	for i, p := range pairs {
		price := PriceCHZPerToken(p, scan.NativeToken)

		priceBigFloat := big.NewFloat(price)
		fmt.Println("FLOAT64", fmt.Sprintf("%f", price))

		fmt.Println("PRICE", priceBigFloat, "PAIR", p.Pair)
		if price <= 0 {
			continue
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

	if buyIndex == -1 || sellIndex == -1 || buyIndex == sellIndex || maxPrice <= minPrice {
		return big.NewInt(0), models.Pair{}, models.Pair{}, fmt.Errorf("alım ve satım için uygun pool bulunamadı veya karlı değil")
	}

	buy := pairs[buyIndex]
	sell := pairs[sellIndex]

	buyRbIn, buyRbOut := getReservesForCHZ(buy, scan.NativeToken.Hex(), true)
	sellRsIn, sellRsOut := getReservesForCHZ(sell, scan.NativeToken.Hex(), false)

	borrow := ComputeBorrowAmount(buyRbIn, buyRbOut, sellRsIn, sellRsOut)

	// Limit: rezervlerin %5'inden fazla borrow edilmesin
	buyLimit := new(big.Int).Div(buyRbIn, big.NewInt(20))    // %5
	sellLimit := new(big.Int).Div(sellRsOut, big.NewInt(20)) // %5
	if sellLimit.Cmp(buyLimit) < 0 {
		if borrow.Cmp(sellLimit) > 0 {
			borrow = sellLimit
		}
	} else {
		if borrow.Cmp(buyLimit) > 0 {
			borrow = buyLimit
		}
	}

	return borrow, buy, sell, nil
}

// CalculateProfit, borrow edilen miktarla arbitraj karını hesaplar.
func CalculateProfit(borrow *big.Int, buy models.Pair, sell models.Pair, chzAddress string) *big.Int {
	buyRbIn, buyRbOut := getReservesForCHZ(buy, chzAddress, true)
	sellRsIn, sellRsOut := getReservesForCHZ(sell, chzAddress, false)

	if borrow.Cmp(big.NewInt(0)) <= 0 {
		return big.NewInt(0)
	}

	buyAmountOut := getAmountOut(borrow, buyRbIn, buyRbOut)          // Token miktarı alındı
	sellAmountOut := getAmountOut(buyAmountOut, sellRsIn, sellRsOut) // Token satılarak CHZ elde edildi

	return new(big.Int).Sub(sellAmountOut, borrow) // Net kar (pozitif olmalı)
}

// getAmountOut, Uniswap V2 formülüne göre output miktarını hesaplar.
func getAmountOut(amountIn, reserveIn, reserveOut *big.Int) *big.Int {
	if amountIn.Cmp(big.NewInt(0)) <= 0 {
		return big.NewInt(0)
	}
	amountInWithFee := new(big.Int).Mul(amountIn, feeMul)
	numerator := new(big.Int).Mul(amountInWithFee, reserveOut)
	denominator := new(big.Int).Add(new(big.Int).Mul(reserveIn, feeDiv), amountInWithFee)
	if denominator.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}
	return numerator.Div(numerator, denominator)
}

// TestComputeOptimalBorrow, örnek pair’lerle borrow ve kar hesaplar.
func TestComputeOptimalBorrow() {
	scan := models.ScanParams{
		NativeToken: *utils.AddressFromHex("0x677f7e16c7dd57be1d4c8ad1244883214953dc47"), // CHZ token adresi
	}

	pairs := []models.Pair{
		{
			BaseReserve:  big.NewInt(500000), // Daha az CHZ, alım havuzu
			QuoteReserve: big.NewInt(1500000000),
			Base:         scan.NativeToken.Hex(),
			Quote:        "0x60f397acbcfb8f4e3234c659a3e10867e6fa6b67",
			Pair:         "BuyPair",
		},
		{
			BaseReserve:  big.NewInt(700000000),
			QuoteReserve: big.NewInt(1500000), // Daha fazla CHZ, satış havuzu
			Base:         "0x60f397acbcfb8f4e3234c659a3e10867e6fa6b67",
			Quote:        scan.NativeToken.Hex(),
			Pair:         "SellPair",
		},
	}
	borrow, buy, sell, err := ComputeOptimalBorrowWithCHZOutput(scan, pairs)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Borrow:", borrow.String())
	fmt.Println("Buy Pair:", buy.Pair)
	fmt.Println("Sell Pair:", sell.Pair)

	profit := CalculateProfit(borrow, buy, sell, scan.NativeToken.Hex())
	fmt.Println("Profit:", profit.String())
}

// FlashSwap, flash swap fırsatını hesaplayıp detayları yazdırır.
func FlashSwap(scan models.ScanParams, pairs []models.Pair) {

	borrow, buyPair, sellPair, err := ComputeOptimalBorrowWithCHZOutput(scan, pairs)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	fmt.Println("Borrow Amount 1 ", borrow, buyPair.Pair, sellPair.Pair)

	chzAddress := scan.NativeToken.Hex()
	buyRbIn, _ := getReservesForCHZ(buyPair, chzAddress, true)

	// Borrow limiti: rezervin %1'i (daha küçük tut)
	maxBorrow := new(big.Int).Div(buyRbIn, big.NewInt(100)) // %1 limit
	if borrow.Cmp(maxBorrow) > 0 {
		borrow = maxBorrow
	}

	if borrow.Cmp(big.NewInt(0)) <= 0 {
		fmt.Println("Borrow is Zero:", borrow)
		return
	}

	profit := CalculateProfit(borrow, buyPair, sellPair, chzAddress)
	if profit.Cmp(big.NewInt(0)) <= 0 {
		fmt.Println("Profit is Zero or Negative:", profit)
		return
	}

	fmt.Printf("FROM %s -> TO %s, Borrow: %s, Profit: %s\n", buyPair.Pair, sellPair.Pair, borrow.String(), profit.String())
	fmt.Printf("Buy Pair Reserves: Base=%s, Quote=%s\n", buyPair.BaseReserve.String(), buyPair.QuoteReserve.String())
	fmt.Printf("Sell Pair Reserves: Base=%s, Quote=%s\n", sellPair.BaseReserve.String(), sellPair.QuoteReserve.String())

	buyRbIn, buyRbOut := getReservesForCHZ(buyPair, chzAddress, true)
	sellRsIn, sellRsOut := getReservesForCHZ(sellPair, chzAddress, false)

	buyAmountOut := getAmountOut(borrow, buyRbIn, buyRbOut)
	sellAmountOut := getAmountOut(buyAmountOut, sellRsIn, sellRsOut)

	fmt.Println("Borrow:", borrow.String())
	fmt.Println("Buy Amount Out (Token):", buyAmountOut.String())
	fmt.Println("Sell Amount Out (CHZ):", sellAmountOut.String())
	fmt.Println("Profit (sellAmountOut - borrow):", profit.String())
}
