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
	DexProtocolV1 DexProtocol = "V1"
	DexProtocolV2 DexProtocol = "V2"
	DexProtocolV3 DexProtocol = "V3"
)

type ChainID int64

const (
	Chiliz    ChainID = 88888
	Ethereum  ChainID = 1
	Avalanche ChainID = 43114
	BSC       ChainID = 56
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
	Ethereum: {
		Explorer: "https://etherscan.io",
		RPC:      "https://eth.drpc.org",
		ChainID:  1,
	},
	Avalanche: {
		Explorer: "https://snowscan.xyz",
		RPC:      "https://api.avax.network/ext/bc/C/rpc",
		ChainID:  43114,
	},
	BSC: {
		Explorer: "https://bscscan.com",
		RPC:      "https://bsc-dataseed.bnbchain.org",
		ChainID:  56,
	},
}

type Exchange struct {
	ID   uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name string       `gorm:"type:text"`
	Kind ExchangeKind `gorm:"type:text"`

	Protocol    DexProtocol     `gorm:"type:text"` //v1 / v2 / v3
	RPC         *string         `gorm:"type:text"`
	Router      *common.Address `gorm:"type:bytea"`
	Factory     *common.Address `gorm:"type:bytea"`
	NativeToken *common.Address `gorm:"type:bytea"`
	Multicall3  *common.Address `gorm:"type:bytea"`
	StablePair  *common.Address `gorm:"type:bytea"`

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
