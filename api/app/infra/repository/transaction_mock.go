package repository

import (
	"context"
)

type mockTransactionRepository struct{}

func NewMockTransactionRepository() *mockTransactionRepository {
	return &mockTransactionRepository{}
}

// WithInTx transactionを貼らずにfunctionだけ実行する
func (m mockTransactionRepository) WithInTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return fn(ctx)
}
