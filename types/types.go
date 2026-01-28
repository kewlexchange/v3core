package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ArbResult struct {
	Exists bool
	Side   bool
	Route  string
	Dx     *big.Int
	Profit *big.Int
	Mid    *big.Int
	Out    *big.Int

	Borrow common.Address
	Repay  common.Address
	Path   []common.Address
}
