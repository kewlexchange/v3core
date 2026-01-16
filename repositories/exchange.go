package repositories

import (
	"context"
	"core/models"

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

func (r *ExchangeRepository) Create(context context.Context, exchangeInfo models.Exchange) error {
	return nil
}

func (r *ExchangeRepository) UpsertPair(context context.Context, exchange models.Exchange) error {
	return nil
}
