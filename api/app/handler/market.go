package handler

import (
	dbRepo "api/app/infra/repository"
	"api/app/oas"
	"api/app/shared/lox"
	"api/app/usecase"
	"context"
	"log/slog"
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

	marketGetRes := oas.MarketGetOKApplicationJSON{}
	for _, market := range markets {
		marketGetRes = append(marketGetRes, oas.MarketGetOKItem{
			ID:          market.ID,
			Name:        market.Name,
			Description: market.Description,
		})
	}

	return &marketGetRes, nil
}
