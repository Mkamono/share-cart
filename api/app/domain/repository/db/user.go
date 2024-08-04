package db

import (
	dbEntity "api/app/domain/entity/db"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *dbEntity.User) error
}
