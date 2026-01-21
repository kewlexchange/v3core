package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Pair struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"  json:"-"`

	ExchangeID uuid.UUID `gorm:"type:uuid;not null"  json:"-"`
	Exchange   Exchange  `gorm:"foreignKey:ExchangeID" json:"-"`

	BaseCurrencyID uuid.UUID `gorm:"type:uuid;not null"  json:"-"`
	BaseCurrency   Currency  `gorm:"foreignKey:BaseCurrencyID"`

	QuoteCurrencyID uuid.UUID `gorm:"type:uuid;not null"  json:"-"`
	QuoteCurrency   Currency  `gorm:"foreignKey:QuoteCurrencyID"`

	Pair  string `gorm:"type:text"`
	Base  string `gorm:"type:text"`
	Quote string `gorm:"type:text"`

	// BASE ASSETIN QUOTEy ye gore FIYATI
	BasePrice *decimal.Decimal `gorm:"type:numeric"`
	// QUOTE ASSETIN BASEYE GORE FIYATI
	QuotePrice *decimal.Decimal `gorm:"type:numeric"`

	// BASE ASSET'in USD TURUNDEN FIYATI
	BasePriceUSD *decimal.Decimal `gorm:"type:numeric"`
	// QUOTE ASSET'in USD TURUNDEN FIYATI
	QuotePriceUSD *decimal.Decimal `gorm:"type:numeric"`

	//BASE ASSETIN CHZ TURUNDEN  FIYATI
	BasePriceNative *decimal.Decimal `gorm:"type:numeric"`
	// QUOTE ASSETIN CHZ TURUNDEN FIYATI
	QuotePriceNative *decimal.Decimal `gorm:"type:numeric"`

	BaseReserve  *string `gorm:"type:text"`
	QuoteReserve *string `gorm:"type:text"`

	BaseDecimals  *string `gorm:"type:text"`
	QuoteDecimals *string `gorm:"type:text"`

	Symbol      string  `gorm:"type:text"`
	DisplayName *string `gorm:"type:text"`

	// Grafik veya UI için
	Logo *string `gorm:"type:text"`

	IsEnabled bool `gorm:"default:true"`
}

func (Pair) TableName() string {
	return "pairs"
}
