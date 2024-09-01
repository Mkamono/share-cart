package handler

import (
	dbRepo "api/app/infra/repository"
	"api/app/oas"
	"api/app/usecase"
	"context"
	"log/slog"
)

func (h *handler) MarketGet(ctx context.Context) ([]oas.Market, error) {
	marketRepo := dbRepo.NewMarketRepository(h.db)
	getMarketsUsecase := usecase.NewGetMarketsUsecase(marketRepo)
	markets, err := getMarketsUsecase.Run(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get markets", "error", err)
		return nil, err
	}
	slog.InfoContext(ctx, "Get markets success")

	oasMarkets := make([]oas.Market, len(markets))
	for i, m := range markets {
		oasMarkets[i] = oas.Market{
			ID:          m.ID,
			Name:        m.Name,
			Description: m.Description,
		}
	}
	return oasMarkets, nil
}
