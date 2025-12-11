package db

import "github.com/google/uuid"

type Currency struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Symbol   string    `gorm:"type:text;not null"` // Global sembol (ETH, BTC, USDT)
	Name     string    `gorm:"type:text"`
	Decimals int32     `gorm:"type:int"`

	Logo       *string `gorm:"type:text"`
	Background *string `gorm:"type:text"`
	Foreground *string `gorm:"type:text"`

	IsEnabled bool `gorm:"default:true"`
}

func (Currency) TableName() string {
	return "currencies"
}
