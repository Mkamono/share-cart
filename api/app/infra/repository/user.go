package repository

import (
	"api/app/domain/entity/db"
	dbRepo "api/app/domain/repository/db"
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ dbRepo.UserRepository = (*userRepository)(nil)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(ctx context.Context, user *db.User) error {
	return user.Insert(ctx, u.db, boil.Infer())
}
