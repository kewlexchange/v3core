package scanner

import (
	"core/models"
	uniswapSDKEntities "core/sdk/uniswap/entities"
	"math/big"

	uniswapSDKConstants "core/sdk/uniswap/constants"
	"fmt"

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

func OptimalInputExact(R1, R2, R3, R4 *big.Int) *big.Int {
	feeNum := big.NewInt(997)
	feeDen := big.NewInt(1000)

	feeNumSq := new(big.Int).Mul(feeNum, feeNum)
	feeDenSq := new(big.Int).Mul(feeDen, feeDen)
	feeNumDen := new(big.Int).Mul(feeNum, feeDen)

	// term1 = R1 * R2 * R3 * R4 * feeNumSq * feeDenSq
	term1 := new(big.Int).Mul(R1, R2)
	term1.Mul(term1, R3)
	term1.Mul(term1, R4)
	term1.Mul(term1, feeNumSq)
	term1.Mul(term1, feeDenSq)

	sqrtTerm := new(big.Int).Sqrt(term1)
	bVal := new(big.Int).Mul(R1, R3)
	bVal.Mul(bVal, feeDenSq)

	if sqrtTerm.Cmp(bVal) <= 0 {
		return big.NewInt(0)
	}

	numerator := new(big.Int).Sub(sqrtTerm, bVal)
	denomPart1 := new(big.Int).Mul(feeNumDen, R3)
	denomPart2 := new(big.Int).Mul(feeNumSq, R2)
	denominator := new(big.Int).Add(denomPart1, denomPart2)

	return new(big.Int).Div(numerator, denominator)
}

func v2AmountOut(amountIn, resIn, resOut *big.Int) *big.Int {
	if amountIn.Cmp(big.NewInt(0)) <= 0 {
		return big.NewInt(0)
	}
	feeNum := big.NewInt(997)
	feeDen := big.NewInt(1000)
	amountInWithFee := new(big.Int).Mul(amountIn, feeNum)
	numerator := new(big.Int).Mul(amountInWithFee, resOut)
	denominator := new(big.Int).Add(new(big.Int).Mul(resIn, feeDen), amountInWithFee)
	return new(big.Int).Div(numerator, denominator)
}

// Python'daki arb_best_two_pools fonksiyonunun Go karşılığı
func ArbBestTwoPools(A_rx, A_ry, B_rx, B_ry *big.Int) {
	// 1. Yön: X -> A -> Y -> B -> X (A'dan al, B'de sat)
	// Pool A: In=X, Out=Y | Pool B: In=Y, Out=X
	opt1 := OptimalInputExact(A_rx, A_ry, B_ry, B_rx)
	profit1 := big.NewInt(0)
	if opt1.Cmp(big.NewInt(0)) > 0 {
		dy := v2AmountOut(opt1, A_rx, A_ry)
		xOut := v2AmountOut(dy, B_ry, B_rx)
		profit1.Sub(xOut, opt1)
	}

	// 2. Yön: X -> B -> Y -> A -> X (B'den al, A'da sat)
	// Pool B: In=X, Out=Y | Pool A: In=Y, Out=X
	opt2 := OptimalInputExact(B_rx, B_ry, A_ry, A_rx)
	profit2 := big.NewInt(0)
	if opt2.Cmp(big.NewInt(0)) > 0 {
		dy := v2AmountOut(opt2, B_rx, B_ry)
		xOut := v2AmountOut(dy, A_ry, A_rx)
		profit2.Sub(xOut, opt2)
	}

	// Karşılaştırma
	if profit1.Cmp(big.NewInt(0)) <= 0 && profit2.Cmp(big.NewInt(0)) <= 0 {
		fmt.Println("İki yönde de kârlı fırsat yok.")
		return
	}

	if profit1.Cmp(profit2) >= 0 {
		fmt.Printf("En iyi rota: A -> B | Giriş: %s | Kâr: %s\n", opt1.String(), profit1.String())
	} else {
		fmt.Printf("En iyi rota: B -> A | Giriş: %s | Kâr: %s\n", opt2.String(), profit2.String())
	}
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
		fmt.Println("Token0Price", pair.Token0Price().ToSignificant(8), price.ToSignificant(8), "TOKEN0:", pair.Token0().Hex())
		fmt.Println("Token1Price", pair.Token1Price().ToSignificant(8), price.ToSignificant(8), "TOKEN1", pair.Token1().Hex())

	}

	fmt.Println("CHEAP PRICE", cheapPrice.ToSignificant(8), cheapPair.Address.Hex())
	fmt.Println("EXPENSIVE PRICE", expensivePrice.ToSignificant(8), expensivePair.Address.Hex())

	var R1, R2, R3, R4 *big.Int

	// cheap
	if cheapPair.Token0().Address == scan.NativeToken {
		R1 = cheapPair.Reserve0().Quotient() // CHZ
		R2 = cheapPair.Reserve1().Quotient() // PEPPER
	} else {
		R1 = cheapPair.Reserve1().Quotient()
		R2 = cheapPair.Reserve0().Quotient()
	}

	// expensive
	if expensivePair.Token0().Address == scan.NativeToken {
		R4 = expensivePair.Reserve0().Quotient() // CHZ
		R3 = expensivePair.Reserve1().Quotient() // PEPPER
	} else {
		R4 = expensivePair.Reserve1().Quotient()
		R3 = expensivePair.Reserve0().Quotient()
	}

	fmt.Println(R1, R2, R3, R4)

	decimals := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

	A_rx := new(big.Int).Mul(big.NewInt(1000), decimals)
	A_ry := new(big.Int).Mul(big.NewInt(2_000_000), decimals)
	B_rx := new(big.Int).Mul(big.NewInt(1200), decimals)
	B_ry := new(big.Int).Mul(big.NewInt(2_000_000), decimals)

	ArbBestTwoPools(A_rx, A_ry, B_rx, B_ry)

}
