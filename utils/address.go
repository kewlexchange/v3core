package utils

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

func AddressFromHex(hex string) *common.Address {
	if hex == "" {
		return nil
	}
	a := common.HexToAddress(hex)
	return &a
}

func StringPtr(s string) *string {
	return &s
}

func BigIntToStringPtr(b *big.Int) *string {
	if b == nil {
		return nil
	}
	str := b.String()
	return &str
}

func StringPtrToBigInt(s *string) *big.Int {
	if s == nil {
		return nil
	}
	bi := new(big.Int)
	bi.SetString(*s, 10)
	return bi
}

// big.Int -> decimal.Decimal
func BigIntToDecimalPtr(b *big.Int) *decimal.Decimal {
	if b == nil {
		return nil
	}
	d := decimal.NewFromBigInt(b, 0)
	return &d
}

// decimal.Decimal -> big.Int
func DecimalToBigInt(d *decimal.Decimal) *big.Int {
	if d == nil {
		return nil
	}
	return d.BigInt()
}

func ParseUnits(amount string, decimals *big.Int) (*big.Int, error) {
	f, ok := new(big.Float).SetString(amount)
	if !ok {
		return nil, fmt.Errorf("invalid amount")
	}

	base := new(big.Int).Exp(big.NewInt(10), decimals, nil) // 10^decimals
	baseFloat := new(big.Float).SetInt(base)

	f.Mul(f, baseFloat)

	i := new(big.Int)
	f.Int(i)

	return i, nil
}

func FormatUnits(amount *big.Int, decimals *big.Int) string {
	f := new(big.Float).SetInt(amount)
	base := new(big.Int).Exp(big.NewInt(10), decimals, nil)
	baseFloat := new(big.Float).SetInt(base)

	f.Quo(f, baseFloat)

	// Ondalık basamak sayısını int olarak almak için:
	decInt := int(decimals.Int64())

	return f.Text('f', decInt)
}
