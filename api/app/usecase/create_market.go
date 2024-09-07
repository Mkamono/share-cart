package usecase

import (
	dbEntity "api/app/domain/entity/db"
	dbRepo "api/app/domain/repository/db"
	"api/app/oas"
	"log/slog"

	"context"
)

type CreateMarket interface {
	Run(ctx context.Context, market *oas.MarketPostReq) (*oas.Market, error)
}

var _ CreateMarket = (*createMarket)(nil)

type createMarket struct {
	marketRepo dbRepo.MarketRepository
}

func NewCreateMarketUsecase(
	marketRepo dbRepo.MarketRepository,
) CreateMarket {
	return &createMarket{
		marketRepo: marketRepo,
	}
}

func (u *createMarket) Run(ctx context.Context, req *oas.MarketPostReq) (*oas.Market, error) {
	market := &dbEntity.Market{}
	market.Name = req.Name
	market.Description = req.Description

	m, err := u.marketRepo.Create(ctx, market)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create market", "error", err)
		return nil, err
	}

	return &oas.Market{
		ID:          m.ID.String(),
		Name:        m.Name,
		Description: m.Description,
	}, nil
}
