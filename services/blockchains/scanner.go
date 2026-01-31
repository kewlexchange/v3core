package blockchains

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

type Scanner interface {
	Start(ctx context.Context) error
	Stop() error
	IsSwapMethod(tx *types.Transaction) bool
}
