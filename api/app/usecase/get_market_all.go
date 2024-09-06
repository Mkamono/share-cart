package usecase

import (
	dbEntity "api/app/domain/entity/db"
	dbRepo "api/app/domain/repository/db"
	"api/app/oas"

	"context"
	"log/slog"

	"github.com/samber/lo"
)

type GetMarketAllUsecase interface {
	Run(ctx context.Context) ([]*oas.Market, error)
}

var _ GetMarketAllUsecase = (*getMarketAllUsecase)(nil)

type getMarketAllUsecase struct {
	marketRepo      dbRepo.MarketRepository
	marketImageRepo dbRepo.MarketImageRepository
}

func NewGetMarketAllUsecase(
	marketRepo dbRepo.MarketRepository,
	marketImageRepo dbRepo.MarketImageRepository,
) GetMarketAllUsecase {
	return &getMarketAllUsecase{
		marketRepo:      marketRepo,
		marketImageRepo: marketImageRepo,
	}
}

func (u *getMarketAllUsecase) Run(ctx context.Context) ([]*oas.Market, error) {
	markets, err := u.marketRepo.GetAll(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get markets", "error", err)
		return nil, err
	}

	if len(markets) == 0 {
		return []*oas.Market{}, nil
	}

	marketImages, err := u.marketImageRepo.GetAllByMarketIDs(ctx, lo.Map(markets, func(m *dbEntity.Market, _ int) string {
		return m.ID
	}))
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get market images", "error", err)
		return nil, err
	}

	slog.InfoContext(ctx, "Success to get markets")

	oasMarkets := lo.Map(markets, func(m *dbEntity.Market, _ int) *oas.Market {
		images := lo.Filter(marketImages, func(i *dbEntity.MarketImage, _ int) bool {
			return i.MarketID == m.ID
		})

		return &oas.Market{
			ID:          m.ID,
			Name:        m.Name,
			Description: m.Description,
			Images: lo.Map(images, func(i *dbEntity.MarketImage, _ int) string {
				return i.ID
			}),
		}
	})

	return oasMarkets, nil
}
