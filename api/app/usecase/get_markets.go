package usecase

import (
	dbEntity "api/app/domain/entity/db"
	dbRepo "api/app/domain/repository/db"
	"api/app/shared/lox"

	"context"
	"log/slog"
)

type GetMarketsUsecase interface {
	Run(ctx context.Context) ([]*dbEntity.Market, error)
}

var _ GetMarketsUsecase = (*getMarketsUsecase)(nil)

type getMarketsUsecase struct {
	marketRepo dbRepo.MarketRepository
}

func NewGetMarketsUsecase(
	marketRepo dbRepo.MarketRepository,
) GetMarketsUsecase {
	return &getMarketsUsecase{
		marketRepo: marketRepo,
	}
}

func (u *getMarketsUsecase) Run(ctx context.Context) ([]*dbEntity.Market, error) {
	markets, err := u.marketRepo.GetAll(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Repo: Failed to get markets", "error", err)
		return nil, err
	}

	slog.InfoContext(ctx, "Repo: Success to get markets", "markets", lox.FromSlicePtr(markets))

	return markets, nil
}
