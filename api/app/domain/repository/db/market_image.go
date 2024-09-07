package db

import (
	dbEntity "api/app/domain/entity/db"
	"context"
)

type MarketImageRepository interface {
	GetAllByMarketIDs(ctx context.Context, marketIDs []string) ([]*dbEntity.MarketImage, error)
	DeleteByMarketID(ctx context.Context, marketID string) error
}
