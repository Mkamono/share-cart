package handler

import (
	dbRepo "api/app/infra/repository"
	"api/app/oas"
	"api/app/shared/lox"
	"api/app/usecase"
	"context"
	"log/slog"

	"github.com/samber/lo"
)

func (h *handler) MarketGet(ctx context.Context) (oas.MarketGetRes, error) {
	marketRepo := dbRepo.NewMarketRepository(h.db)
	getMarketsUsecase := usecase.NewGetMarketsUsecase(marketRepo)
	markets, err := getMarketsUsecase.Run(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Usecase: Failed to get markets", "error", err)
		return nil, err
	}
	slog.InfoContext(ctx, "Usecase: Get markets success", "markets", lox.FromSlicePtr(markets))

	oasMarkets := make([]oas.MarketGetOKItem, len(markets))
	for i, m := range markets {
		oasMarkets[i] = oas.MarketGetOKItem{
			ID:          m.ID,
			Name:        m.Name,
			Description: m.Description,
		}
	}

	return lo.ToPtr(append(oas.MarketGetOKApplicationJSON{}, oasMarkets...)), nil
}
