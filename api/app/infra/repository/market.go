package repository

import (
	"api/app/domain/entity/db"
	dbRepo "api/app/domain/repository/db"
	"api/app/infra/boiler"
	"context"
	"database/sql"

	"github.com/samber/lo"
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
		return nil, err
	}

	markets := lo.Map(boilerMarkets, func(x *boiler.Market, _ int) *db.Market {
		return &db.Market{
			Market: *x,
		}
	})

	return markets, nil
}
