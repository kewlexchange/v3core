package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type ExchangeKind string

const (
	ExchangeKindCEX ExchangeKind = "CEX"
	ExchangeKindDEX ExchangeKind = "DEX"
)

type DexProtocol string

const (
	DexProtocolV2 DexProtocol = "V2"
	DexProtocolV3 DexProtocol = "V3"
)

type ChainID int64

const (
	Chiliz ChainID = 88888
)

type ChainInfo struct {
	Explorer string
	RPC      string
	ChainID  int64
}

var ChainInfoMap = map[ChainID]ChainInfo{
	Chiliz: {
		Explorer: "https://chiliscan.io",
		RPC:      "https://rpc.chiliz.com",
		ChainID:  88888,
	},
}

type Exchange struct {
	ID   uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name string       `gorm:"type:text"`
	Kind ExchangeKind `gorm:"type:text"`

	Protocol    DexProtocol     `gorm:"type:text"` // V2 / V3 (sadece DEX)
	RPC         *string         `gorm:"type:text"`
	Router      *common.Address `gorm:"type:bytea"`
	Factory     *common.Address `gorm:"type:bytea"`
	NativeToken *common.Address `gorm:"type:bytea"`
	Multicall3  *common.Address `gorm:"type:bytea"`

	ChainID    *int64  `gorm:"type:bigint"`
	Explorer   *string `gorm:"type:text"`
	Background *string `gorm:"type:text"`
	Foreground *string `gorm:"type:text"`
	Logo       *string `gorm:"type:text"`

	IsEnabled bool `gorm:"default:true"`
}

func (Exchange) TableName() string {
	return "exchanges"
}
