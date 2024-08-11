package db

import (
	dbEntity "api/app/domain/entity/db"
	"context"
)

type MarketRepository interface {
	GetAll(ctx context.Context) ([]*dbEntity.Market, error)
}
