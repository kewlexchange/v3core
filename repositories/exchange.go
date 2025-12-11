package repositories

import (
	"context"
	"core/models/db"

	"gorm.io/gorm"
)

type ExchangeRepository struct {
	db *gorm.DB
}

func (r *ExchangeRepository) DB() *gorm.DB {
	return r.db
}

func NewExchangeRepository(db *gorm.DB) *ExchangeRepository {
	return &ExchangeRepository{db: db}
}

func (r *ExchangeRepository) Create(context context.Context, exchangeInfo db.Exchange) error {
	return nil
}

func (r *ExchangeRepository) UpsertPair(context context.Context, exchange db.Exchange) error {
	return nil
}
