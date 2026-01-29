package models

import "github.com/ethereum/go-ethereum/common"

type Hop struct {
	TradingPair *TradingPair
	Pair        common.Address
	OutputToken common.Address
}

type Cycle struct {
	InputToken common.Address
	Hops       []Hop
}
