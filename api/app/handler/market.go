package handler

import (
	dbRepo "api/app/infra/repository"
	"api/app/oas"
	"api/app/shared/lox"
	"api/app/usecase"
	"context"
	"log/slog"
)

type MarketDTO struct {
	oas.Market
}

func (h *handler) MarketGet(ctx context.Context) ([]oas.Market, error) {
	marketRepo := dbRepo.NewMarketRepository(h.db)
	marketImageRepo := dbRepo.NewMarketImageRepository(h.db)
	getMarketsUsecase := usecase.NewGetMarketsUsecase(marketRepo, marketImageRepo)

	oasMarkets, err := getMarketsUsecase.Run(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get markets", "error", err)
		return nil, err
	}
	slog.InfoContext(ctx, "Get markets success")

	return lox.FromSlicePtr(oasMarkets), nil
}
