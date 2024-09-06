package handler

import (
	dbRepo "api/app/infra/repository"
	"api/app/oas"
	"api/app/shared/lox"
	"api/app/usecase"
	"context"
	"log/slog"
)

func (h *handler) MarketGet(ctx context.Context) ([]oas.Market, error) {
	marketRepo := dbRepo.NewMarketRepository(h.db)
	marketImageRepo := dbRepo.NewMarketImageRepository(h.db)
	getMarketAllUsecase := usecase.NewGetMarketAllUsecase(marketRepo, marketImageRepo)

	oasMarkets, err := getMarketAllUsecase.Run(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get markets", "error", err)
		return nil, err
	}
	slog.InfoContext(ctx, "Get markets success")

	return lox.FromSlicePtr(oasMarkets), nil
}
