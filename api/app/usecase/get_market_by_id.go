package usecase

import (
	dbEntity "api/app/domain/entity/db"
	dbRepo "api/app/domain/repository/db"
	"api/app/oas"
	"log/slog"

	"context"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type GetMarketByID interface {
	Run(ctx context.Context, id uuid.UUID) (*oas.Market, error)
}

var _ GetMarketByID = (*getMarketByID)(nil)

type getMarketByID struct {
	marketRepo      dbRepo.MarketRepository
	marketImageRepo dbRepo.MarketImageRepository
}

func NewGetMarketByIDUsecase(
	marketRepo dbRepo.MarketRepository,
	marketImageRepo dbRepo.MarketImageRepository,
) GetMarketByID {
	return &getMarketByID{
		marketRepo:      marketRepo,
		marketImageRepo: marketImageRepo,
	}
}

func (u *getMarketByID) Run(ctx context.Context, id uuid.UUID) (*oas.Market, error) {
	market, err := u.marketRepo.GetByID(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to retrieve market", "error", err)
		return nil, err
	}

	marketImages, err := u.marketImageRepo.GetAllByMarketIDs(ctx, []uuid.UUID{market.ID})
	if err != nil {
		slog.ErrorContext(ctx, "failed to retrieve market images", "error", err)
		return nil, err
	}

	return &oas.Market{
		ID:          market.ID.String(),
		Name:        market.Name,
		Description: market.Description,
		Images: lo.Map(marketImages, func(i *dbEntity.MarketImage, _ int) string {
			return i.ImageID
		}),
	}, nil
}
