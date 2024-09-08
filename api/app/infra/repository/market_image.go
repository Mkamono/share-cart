package repository

import (
	"api/app/domain/entity/db"
	dbRepo "api/app/domain/repository/db"
	"api/app/infra/boiler"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"strings"

	"github.com/google/uuid"
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

func (m *marketImageRepository) GetAllByMarketIDs(ctx context.Context, marketIDs uuid.UUIDs) ([]*db.MarketImage, error) {
	IDs := lo.Map(marketIDs, func(x uuid.UUID, _ int) string {
		return x.String()
	})

	boilerMarketImages, err := boiler.MarketImages(
		boiler.MarketImageWhere.MarketID.IN(IDs),
	).All(ctx, m.db)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("failed to get market images. marketIDs = %s", strings.Join(IDs, ", ")), "error", err)
		return nil, err
	}

	marketImages := lo.Map(boilerMarketImages, func(x *boiler.MarketImage, _ int) *db.MarketImage {
		return &db.MarketImage{
			MarketImage: *x,
		}
	})

	return marketImages, nil
}

func (m *marketImageRepository) DeleteByMarketID(ctx context.Context, marketID uuid.UUID) error {
	_, err := boiler.MarketImages(
		boiler.MarketImageWhere.MarketID.EQ(marketID.String()),
	).DeleteAll(ctx, m.db)
	if err != nil {
		slog.ErrorContext(ctx, "failed to delete market images", "error", err)
		return err
	}

	return nil
}
