package db

import (
	"github.com/google/uuid"
)

type ExchangeKind string

const (
	ExchangeKindCEX ExchangeKind = "CEX"
	ExchangeKindDEX ExchangeKind = "DEX"
)

type Exchange struct {
	ID   uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name string       `gorm:"type:text"`
	Kind ExchangeKind `gorm:"type:text"`

	RPC        *string `gorm:"type:text"`
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
