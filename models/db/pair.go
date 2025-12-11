package db

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Pair struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`

	ExchangeID uuid.UUID `gorm:"type:uuid;not null"`
	Exchange   Exchange  `gorm:"foreignKey:ExchangeID"`

	BaseCurrencyID uuid.UUID `gorm:"type:uuid;not null"`
	BaseCurrency   Currency  `gorm:"foreignKey:BaseCurrencyID"`

	QuoteCurrencyID uuid.UUID `gorm:"type:uuid;not null"`
	QuoteCurrency   Currency  `gorm:"foreignKey:QuoteCurrencyID"`

	Base  string `gorm:"type:text"`
	Quote string `gorm:"type:text"`

	BasePrice  *decimal.Decimal `gorm:"type:numeric"`
	QuotePrice *decimal.Decimal `gorm:"type:numeric"`

	BaseReserve  *string `gorm:"type:text"`
	QuoteReserve *string `gorm:"type:text"`

	Symbol      string  `gorm:"type:text"`
	DisplayName *string `gorm:"type:text"`

	// Grafik veya UI için
	Logo *string `gorm:"type:text"`

	IsEnabled bool `gorm:"default:true"`
}

func (Pair) TableName() string {
	return "pairs"
}
