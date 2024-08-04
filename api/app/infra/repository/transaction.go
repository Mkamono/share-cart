package repository

import (
	dbRepo "api/app/domain/repository/db"
	"context"
	"database/sql"
)

var _ dbRepo.TransactionRepository = (*transactionRepository)(nil)

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *transactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) WithInTx(ctx context.Context, f func(ctx context.Context) error) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := f(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
