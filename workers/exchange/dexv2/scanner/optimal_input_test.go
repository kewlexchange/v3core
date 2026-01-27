package scanner

import (
	"fmt"
	"math/big"
	"testing"
)

func bi(s string) *big.Int {
	v, _ := new(big.Int).SetString(s, 10)
	return v
}

func TestOptimalInputExact_REAL_RESERVES(t *testing.T) {
	decimals := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

	// Python Örneği:
	// B_rx = 1000, B_ry = 2M
	// A_rx = 1200, A_ry = 2M
	B_rx := new(big.Int).Mul(big.NewInt(1000), decimals)
	B_ry := new(big.Int).Mul(big.NewInt(2_000_000), decimals)
	A_rx := new(big.Int).Mul(big.NewInt(1200), decimals)
	A_ry := new(big.Int).Mul(big.NewInt(2_000_000), decimals)

	res := ArbBestTwoPools(A_rx, A_ry, B_rx, B_ry)

	fmt.Printf("route: %s\n", res.Route)
	fmt.Printf("dx: %s\n", res.Dx.String())
	fmt.Printf("profit: %s\n", res.Profit.String())
	fmt.Printf("mid: %s\n", res.Mid.String())
	fmt.Printf("out: %s\n", res.Out.String())
}

func TestOptimalInputExact_WITH_ARB(t *testing.T) {
	decimals := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

	// Python Örneği:
	// B_rx = 1000, B_ry = 2M
	// A_rx = 1200, A_ry = 2M
	A_rx := new(big.Int).Mul(big.NewInt(1000), decimals)
	A_ry := new(big.Int).Mul(big.NewInt(2_000_000), decimals)
	B_rx := new(big.Int).Mul(big.NewInt(1200), decimals)
	B_ry := new(big.Int).Mul(big.NewInt(2_000_000), decimals)

	res := ArbBestTwoPools(A_rx, A_ry, B_rx, B_ry)

	fmt.Printf("route: %s\n", res.Route)
	fmt.Printf("dx: %s\n", res.Dx.String())
	fmt.Printf("profit: %s\n", res.Profit.String())
	fmt.Printf("mid: %s\n", res.Mid.String())
	fmt.Printf("out: %s\n", res.Out.String())
}
