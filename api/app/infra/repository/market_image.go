package repository

import (
	"api/app/domain/entity/db"
	dbRepo "api/app/domain/repository/db"
	"api/app/infra/boiler"
	"context"
	"database/sql"
	"log/slog"

	"github.com/samber/lo"
)

var _ dbRepo.MarketImageRepository = (*marketImageRepository)(nil)

type marketImageRepository struct {
	db *sql.DB
}

func NewMarketImageRepository(db *sql.DB) *marketImageRepository {
	return &marketImageRepository{
		db: db,
	}
}

func (m *marketImageRepository) GetAllByMarketIDs(ctx context.Context, marketIDs []string) ([]*db.MarketImage, error) {
	boilerMarketImages, err := boiler.MarketImages(
		boiler.MarketImageWhere.MarketID.IN(marketIDs),
	).All(ctx, m.db)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get market images", "error", err)
		return nil, err
	}

	marketImages := lo.Map(boilerMarketImages, func(x *boiler.MarketImage, _ int) *db.MarketImage {
		return &db.MarketImage{
			MarketImage: *x,
		}
	})

	return marketImages, nil
}
