package usecase

import (
	dbEntity "api/app/domain/entity/db"
	dbRepo "api/app/domain/repository/db"
	"api/app/infra/boiler"
	"api/app/oas"

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
	m, err := u.marketRepo.Create(ctx, &dbEntity.Market{
		Market: boiler.Market{
			Name:        req.Name,
			Description: req.Description,
		},
	})
	if err != nil {
		return nil, err
	}

	return &oas.Market{
		ID:          m.ID.String(),
		Name:        m.Name,
		Description: m.Description,
	}, nil
}
