package db

import (
	"context"
)

type TransactionRepository interface {
	WithInTx(ctx context.Context, f func(ctx context.Context) error) error
}
