package usecase

import (
	dbRepo "api/app/domain/repository/db"
	"context"

	"github.com/google/uuid"
)

type DeleteMarketByID interface {
	Run(ctx context.Context, id uuid.UUID) error
}

var _ DeleteMarketByID = (*deleteMarketByID)(nil)

type deleteMarketByID struct {
	txRepo          dbRepo.TransactionRepository
	marketRepo      dbRepo.MarketRepository
	marketImageRepo dbRepo.MarketImageRepository
}

func NewDeleteMarketByIDUsecase(
	txRepo dbRepo.TransactionRepository,
	marketRepo dbRepo.MarketRepository,
	marketImageRepo dbRepo.MarketImageRepository,
) DeleteMarketByID {
	return &deleteMarketByID{
		txRepo:          txRepo,
		marketRepo:      marketRepo,
		marketImageRepo: marketImageRepo,
	}
}

func (u *deleteMarketByID) Run(ctx context.Context, id uuid.UUID) error {
	u.txRepo.WithInTx(ctx, func(ctx context.Context) error {
		err := u.marketRepo.DeleteByID(ctx, id)
		if err != nil {
			return err
		}
		err = u.marketImageRepo.DeleteByMarketID(ctx, id)
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}
