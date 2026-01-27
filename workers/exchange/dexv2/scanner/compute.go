package scanner

import (
	"core/models"
	uniswapSDKConstants "core/sdk/uniswap/constants"
	uniswapSDKEntities "core/sdk/uniswap/entities"
	coreUtils "core/utils"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func GetPrice(nativeCurrency common.Address, pair *uniswapSDKEntities.Pair) *uniswapSDKEntities.Price {
	var price *uniswapSDKEntities.Price
	isCHZToken0 := pair.Token0().Address == nativeCurrency
	isCHZToken1 := pair.Token1().Address == nativeCurrency

	if isCHZToken0 {
		price = pair.Token0Price()
	} else if isCHZToken1 {
		price = pair.Token1Price()
	}
	return price
}

type ArbResult struct {
	Exists bool
	Route  string
	Dx     *big.Int
	Profit *big.Int
	Mid    *big.Int
	Out    *big.Int
}

func OptimalInputExact(rIn1, rOut1, rIn2, rOut2 *big.Int) *big.Int {
	feeNum := big.NewInt(997)
	feeDen := big.NewInt(1000)

	fNumSq := new(big.Int).Mul(feeNum, feeNum)
	fDenSq := new(big.Int).Mul(feeDen, feeDen)
	fNumDen := new(big.Int).Mul(feeNum, feeDen)

	term1 := new(big.Int).Mul(rIn1, rOut1)
	term1.Mul(term1, rIn2)
	term1.Mul(term1, rOut2)
	term1.Mul(term1, fNumSq)
	term1.Mul(term1, fDenSq)

	sqrtTerm := new(big.Int).Sqrt(term1)

	bVal := new(big.Int).Mul(rIn1, rIn2)
	bVal.Mul(bVal, fDenSq)

	if sqrtTerm.Cmp(bVal) <= 0 {
		return big.NewInt(0)
	}

	numerator := new(big.Int).Sub(sqrtTerm, bVal)

	d1 := new(big.Int).Mul(fNumDen, rIn2)
	d2 := new(big.Int).Mul(fNumSq, rOut1)
	denominator := new(big.Int).Add(d1, d2)

	return new(big.Int).Div(numerator, denominator)
}

func v2AmountOut(dx, rx, ry *big.Int) *big.Int {
	if dx.Sign() <= 0 {
		return big.NewInt(0)
	}

	feeNum := big.NewInt(997)
	feeDen := big.NewInt(1000)

	dxEff := new(big.Int).Mul(dx, feeNum)
	num := new(big.Int).Mul(dxEff, ry)
	den := new(big.Int).Add(new(big.Int).Mul(rx, feeDen), dxEff)

	fmt.Println("AmountOut", dx, rx, ry, new(big.Int).Div(num, den))

	return new(big.Int).Div(num, den)
}

func PriceImpact(dx, rx, ry *big.Int) *uniswapSDKEntities.Percent {
	const scale = 1e18

	spotPriceF := new(big.Float).Quo(new(big.Float).SetInt(ry), new(big.Float).SetInt(rx))
	amountOut := v2AmountOut(dx, rx, ry)
	if amountOut.Cmp(big.NewInt(0)) == 0 {
		return uniswapSDKEntities.NewPercent(big.NewInt(0), big.NewInt(1))
	}

	dxF := new(big.Float).SetInt(dx)
	amountOutF := new(big.Float).SetInt(amountOut)
	effectivePriceF := new(big.Float).Quo(amountOutF, dxF)

	diff := new(big.Float).Sub(spotPriceF, effectivePriceF)
	diffAbs := new(big.Float).Abs(diff)
	ratioF := new(big.Float).Quo(diffAbs, spotPriceF)

	scaleF := new(big.Float).SetFloat64(scale)
	scaledRatioF := new(big.Float).Mul(ratioF, scaleF)

	ratioInt := new(big.Int)
	scaledRatioF.Int(ratioInt)

	return uniswapSDKEntities.NewPercent(ratioInt, big.NewInt(scale))
}

func ArbBestTwoPools(A_rx, A_ry, B_rx, B_ry *big.Int) ArbResult {
	maxFracNum := big.NewInt(3)
	maxFracDen := big.NewInt(10)
	dx1 := OptimalInputExact(A_rx, A_ry, B_ry, B_rx)
	limit1 := new(big.Int).Mul(A_rx, maxFracNum)
	limit1.Div(limit1, maxFracDen)
	if dx1.Cmp(limit1) > 0 {
		dx1.Set(limit1)
	}

	p1 := big.NewInt(0)
	var dy1, out1 *big.Int = big.NewInt(0), big.NewInt(0)
	if dx1.Sign() > 0 {
		dy1 = v2AmountOut(dx1, A_rx, A_ry)
		out1 = v2AmountOut(dy1, B_ry, B_rx)
		p1.Sub(out1, dx1)
	}

	dx2 := OptimalInputExact(B_rx, B_ry, A_ry, A_rx)

	limit2 := new(big.Int).Mul(B_rx, maxFracNum)
	limit2.Div(limit2, maxFracDen)
	if dx2.Cmp(limit2) > 0 {
		dx2.Set(limit2)
	}

	p2 := big.NewInt(0)
	var dy2, out2 *big.Int = big.NewInt(0), big.NewInt(0)
	if dx2.Sign() > 0 {
		dy2 = v2AmountOut(dx2, B_rx, B_ry)
		out2 = v2AmountOut(dy2, A_ry, A_rx)
		p2.Sub(out2, dx2)
	}

	if p1.Cmp(p2) >= 0 && p1.Sign() > 0 {
		return ArbResult{Exists: true, Route: "A->B", Dx: dx1, Profit: p1, Mid: dy1, Out: out1}
		//return ArbResult{"A->B", dx1, p1, dy1, out1}
	} else if p2.Sign() > 0 {
		return ArbResult{Exists: true, Route: "B->A", Dx: dx2, Profit: p2, Mid: dy2, Out: out2}
		//return ArbResult{"B->A", dx2, p2, dy2, out2}
	}

	return ArbResult{Exists: false, Route: "", Dx: big.NewInt(0), Profit: big.NewInt(0), Out: big.NewInt(0)}
}

func FlashSwap(scan models.ScanParams, tradingPairs []models.TradingPair) {

	var cheapPair *uniswapSDKEntities.Pair
	var cheapPrice *uniswapSDKEntities.Price
	var expensivePair *uniswapSDKEntities.Pair
	var expensivePrice *uniswapSDKEntities.Price

	for _, tradingPair := range tradingPairs {

		//	baseChainId := int(*tradingPair.BaseCurrency.ChainID)
		//		quoteChainId := int(*tradingPair.QuoteCurrency.ChainID)
		baseChainID := uniswapSDKConstants.ChainID(*tradingPair.BaseCurrency.ChainID)
		quoteChainID := uniswapSDKConstants.ChainID(*tradingPair.QuoteCurrency.ChainID)

		token0, token0Err := uniswapSDKEntities.NewToken(baseChainID, *tradingPair.BaseCurrency.Contract, int(tradingPair.BaseCurrency.Decimals.Int64()), "TOKEN0", "TOKEN0_NAME")
		token1, token1Err := uniswapSDKEntities.NewToken(quoteChainID, *tradingPair.QuoteCurrency.Contract, int(tradingPair.QuoteCurrency.Decimals.Int64()), "TOKEN1", "TOKEN1_NAME")

		if token0Err != nil || token1Err != nil {
			fmt.Println("Token0 ya da Token1 hatali")
			return
		}

		tokenAmount0, tokenAmount0Err := uniswapSDKEntities.NewTokenAmount(token0, tradingPair.BaseReserve)
		tokenAmount1, tokenAmount1Err := uniswapSDKEntities.NewTokenAmount(token1, tradingPair.QuoteReserve)

		if tokenAmount0Err != nil || tokenAmount1Err != nil {
			fmt.Println("tokenAmount0Err ya da tokenAmount1Err hatali")
			return
		}

		pair, pairErr := uniswapSDKEntities.NewPair(tokenAmount0, tokenAmount1, &tradingPair.Pair)
		if pairErr != nil {
			fmt.Println("pairErr hatali")
			return
		}

		price := GetPrice(scan.NativeToken, pair)

		// CHEAP
		if cheapPrice == nil || price.LessThan(cheapPrice.Fraction) {
			cheapPair = pair
			cheapPrice = price
		}

		// EXPENSIVE
		if expensivePrice == nil || price.GreaterThan(expensivePrice.Fraction) {
			expensivePair = pair
			expensivePrice = price
		}

		fmt.Println("Pair", pair.Address)
		fmt.Println("Token0Price", pair.Token0Price().ToSignificant(8), "TOKEN0:", pair.Token0().Hex())
		fmt.Println("Token1Price", pair.Token1Price().ToSignificant(8), "TOKEN1", pair.Token1().Hex())

	}

	fmt.Println("CHEAP PRICE", cheapPrice.ToSignificant(8), cheapPair.Address.Hex())
	fmt.Println("EXPENSIVE PRICE", expensivePrice.ToSignificant(8), expensivePair.Address.Hex())

	var R1, R2, R3, R4 *big.Int

	// cheap
	if cheapPair.Token0().Address == scan.NativeToken {
		R1 = cheapPair.Reserve0().Raw() // CHZ
		R2 = cheapPair.Reserve1().Raw() // PEPPER
	} else {
		R1 = cheapPair.Reserve1().Raw()
		R2 = cheapPair.Reserve0().Raw()
	}

	// expensive
	if expensivePair.Token0().Address == scan.NativeToken {
		R3 = expensivePair.Reserve0().Raw() // Native Token
		R4 = expensivePair.Reserve1().Raw() // Diğer token
	} else {
		R3 = expensivePair.Reserve1().Raw()
		R4 = expensivePair.Reserve0().Raw()
	}

	res := ArbBestTwoPools(R1, R2, R3, R4)
	if !res.Exists {
		fmt.Println("Arbitrage Not Found")
		return
	}
	fmt.Printf("route: %s\n", res.Route)
	fmt.Printf("dx Ether: %s\n", coreUtils.ToEther(res.Dx))
	fmt.Printf("profit Ether: %s\n", coreUtils.ToEther(res.Profit))
	fmt.Printf("mid: %s\n", coreUtils.ToEther(res.Mid))
	fmt.Printf("out: %s\n", coreUtils.ToEther(res.Out))

	priceImpactA := PriceImpact(res.Dx, R1, R2)
	priceImpactB := PriceImpact(res.Dx, R3, R4)

	fmt.Printf("Price Impact on Pool A: %s\n", priceImpactA.ToFixed(2))
	fmt.Printf("Price Impact on Pool B: %s\n", priceImpactB.ToFixed(2))

}
