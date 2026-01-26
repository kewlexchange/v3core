package models

import "github.com/ethereum/go-ethereum/common"

type ScanParams struct {
	Token       string         `gorm:"type:text"`
	NativeToken common.Address `gorm:"type:text"`
	Pairs       []TradingPair
}
