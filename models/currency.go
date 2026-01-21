package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type Currency struct {
	ID         uuid.UUID       `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"  json:"-"`
	Symbol     string          `gorm:"type:text;not null"` // Global sembol (ETH, BTC, USDT)
	Name       string          `gorm:"type:text"  json:"-"`
	Decimals   *string         `gorm:"type:text"  json:"-"`
	Contract   *common.Address `gorm:"type:bytea"`
	ChainID    *int64          `gorm:"type:bigint"  json:"-"`
	Logo       *string         `gorm:"type:text"  json:"-"`
	Background *string         `gorm:"type:text" json:"-" `
	Foreground *string         `gorm:"type:text" json:"-"`
	IsEnabled  bool            `gorm:"default:true"  json:"-"`
}

func (Currency) TableName() string {
	return "currencies"
}
