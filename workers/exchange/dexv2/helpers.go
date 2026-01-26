package dexv2

import "math/big"

func GetAmountOut(amountIn, reserveIn, reserveOut *big.Int) *big.Int {
	if amountIn.Cmp(big.NewInt(0)) <= 0 {
		return big.NewInt(0)
	}
	feeMul := big.NewInt(997)
	feeDiv := big.NewInt(1000)

	amountInWithFee := new(big.Int).Mul(amountIn, feeMul)
	numerator := new(big.Int).Mul(amountInWithFee, reserveOut)
	denominator := new(big.Int).Add(new(big.Int).Mul(reserveIn, feeDiv), amountInWithFee)
	if denominator.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}
	return new(big.Int).Div(numerator, denominator)
}

func GetAmountIn(amountOut, reserveIn, reserveOut *big.Int) *big.Int {
	if amountOut.Cmp(big.NewInt(0)) <= 0 {
		return big.NewInt(0)
	}
	feeMul := big.NewInt(997)
	feeDiv := big.NewInt(1000)

	numerator := new(big.Int).Mul(new(big.Int).Mul(reserveIn, amountOut), feeDiv)
	denominator := new(big.Int).Mul(new(big.Int).Sub(reserveOut, amountOut), feeMul)
	if denominator.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}
	amountIn := new(big.Int).Div(numerator, denominator)
	amountIn.Add(amountIn, big.NewInt(1)) // +1 to round up
	return amountIn
}
