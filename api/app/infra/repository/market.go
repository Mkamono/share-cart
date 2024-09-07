package repository

import (
	"api/app/domain/entity/db"
	dbRepo "api/app/domain/repository/db"
	"api/app/infra/boiler"
	"context"
	"database/sql"
	"log/slog"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ dbRepo.MarketRepository = (*marketRepository)(nil)

type marketRepository struct {
	db *sql.DB
}

func NewMarketRepository(db *sql.DB) *marketRepository {
	return &marketRepository{
		db: db,
	}
}

func (m *marketRepository) GetAll(ctx context.Context) ([]*db.Market, error) {
	boilerMarkets, err := boiler.Markets().All(ctx, m.db)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get markets", "error", err)
		return nil, err
	}

	markets := lo.Map(boilerMarkets, func(x *boiler.Market, _ int) *db.Market {
		return &db.Market{
			Market: *x,
		}
	})

	return markets, nil
}

func (m *marketRepository) Create(ctx context.Context, market *db.Market) (*db.Market, error) {
	market.ID = uuid.New()
	err := market.Insert(ctx, m.db, boil.Infer())
	if err != nil {
		slog.ErrorContext(ctx, "failed to create market", "error", err)
		return nil, err
	}

	return market, nil
}

func (m *marketRepository) GetByID(ctx context.Context, id uuid.UUID) (*db.Market, error) {
	boilerMarket, err := boiler.FindMarket(ctx, m.db, id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get market", "error", err)
		return nil, err
	}

	return &db.Market{
		Market: *boilerMarket,
	}, nil
}

func (m *marketRepository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	boilerMarket, err := boiler.FindMarket(ctx, m.db, id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get market", "error", err)
		return err
	}

	_, err = boilerMarket.Delete(ctx, m.db)
	if err != nil {
		slog.ErrorContext(ctx, "failed to delete market", "error", err)
		return err
	}

	return nil
}
