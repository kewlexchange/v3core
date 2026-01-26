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

	fNum := big.NewInt(997)
	fDen := big.NewInt(1000)

	// t = sqrt(R1 * R2 * R3 * R4)
	t := new(big.Int).Mul(R1, R2)
	t.Mul(t, R3)
	t.Mul(t, R4)

	sqrt := new(big.Int).Sqrt(t)

	// sqrt * f
	sqrt.Mul(sqrt, fNum)
	sqrt.Div(sqrt, fDen)

	b := new(big.Int).Mul(R1, R3)

	fmt.Println("sqrt:", sqrt.String())
	fmt.Println("b:", b.String())

	if sqrt.Cmp(b) <= 0 {
		fmt.Println("NO ARB")
		return big.NewInt(0)
	}

	numerator := new(big.Int).Sub(sqrt, b)

	denominator := new(big.Int).Add(R3, R2)
	denominator.Mul(denominator, fNum)
	denominator.Div(denominator, fDen)

	fmt.Println("numerator:", numerator.String())
	fmt.Println("denominator:", denominator.String())

	return numerator.Div(numerator, denominator)
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

	optimal := OptimalInputExact(R1, R2, R3, R4)

	fmt.Println("OptimalInput", optimal)

}
