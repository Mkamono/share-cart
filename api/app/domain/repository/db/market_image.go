package db

import (
	dbEntity "api/app/domain/entity/db"
	"context"

	"github.com/google/uuid"
)

type MarketImageRepository interface {
	GetAllByMarketIDs(ctx context.Context, marketIDs uuid.UUIDs) ([]*dbEntity.MarketImage, error)
	DeleteByMarketID(ctx context.Context, marketID uuid.UUID) error
}
