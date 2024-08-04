package db

import (
	dbEntity "api/app/domain/entity/db"
	"context"
)

type AuthSubjectRepository interface {
	Create(ctx context.Context, user *dbEntity.AuthSubject) error
}
