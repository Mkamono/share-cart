package handler

import (
	dbRepo "api/app/infra/repository"
	"api/app/oas"
	"api/app/shared/ctxlogger"
	"api/app/shared/lox"
	"api/app/usecase"
	"context"
	"log/slog"

	"github.com/google/uuid"
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

func (h *handler) MarketPost(ctx context.Context, req *oas.MarketPostReq) (*oas.Market, error) {
	ctxlogger.WithValue(ctx, "req", req)

	marketRepo := dbRepo.NewMarketRepository(h.db)
	createMarketUsecase := usecase.NewCreateMarketUsecase(marketRepo)

	market, err := createMarketUsecase.Run(ctx, req)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to create market", "error", err)
		return nil, err
	}

	return &oas.Market{
		ID:          market.ID,
		Name:        market.Name,
		Description: market.Description,
	}, nil
}

func (h *handler) MarketMarketIdDelete(ctx context.Context, params oas.MarketMarketIdDeleteParams) error {
	ctxlogger.WithValue(ctx, "params", params)

	txRepo := dbRepo.NewTransactionRepository(h.db)
	marketRepo := dbRepo.NewMarketRepository(h.db)
	marketImageRepo := dbRepo.NewMarketImageRepository(h.db)
	deleteMarketByIDUsecase := usecase.NewDeleteMarketByIDUsecase(txRepo, marketRepo, marketImageRepo)

	marketID, err := uuid.Parse(params.MarketId)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to parse market id", "error", err)
		return err
	}
	err = deleteMarketByIDUsecase.Run(ctx, marketID)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to delete market", "error", err)
		return err
	}
	slog.InfoContext(ctx, "Delete market success")

	return nil
}

func (h *handler) MarketMarketIdGet(ctx context.Context, params oas.MarketMarketIdGetParams) (*oas.Market, error) {
	ctxlogger.WithValue(ctx, "params", params)

	marketRepo := dbRepo.NewMarketRepository(h.db)
	marketImageRepo := dbRepo.NewMarketImageRepository(h.db)
	getMarketByIDUsecase := usecase.NewGetMarketByIDUsecase(marketRepo, marketImageRepo)

	marketID, err := uuid.Parse(params.MarketId)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to parse market id", "error", err)
		return nil, err
	}
	oasMarket, err := getMarketByIDUsecase.Run(ctx, marketID)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get market", "error", err)
		return nil, err
	}
	slog.InfoContext(ctx, "Get market success")

	return oasMarket, nil
}
