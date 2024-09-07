package db

import (
	dbEntity "api/app/domain/entity/db"
	"context"

	"github.com/google/uuid"
)

type MarketRepository interface {
	GetAll(ctx context.Context) ([]*dbEntity.Market, error)
	Create(ctx context.Context, market *dbEntity.Market) (*dbEntity.Market, error)
	GetByID(ctx context.Context, id uuid.UUID) (*dbEntity.Market, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
}
