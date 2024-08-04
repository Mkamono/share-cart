package repository

import (
	"api/app/domain/entity/db"
	dbRepo "api/app/domain/repository/db"
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var _ dbRepo.AuthSubjectRepository = (*authSubjectRepository)(nil)

type authSubjectRepository struct {
	db *sql.DB
}

func NewAuthSubjectRepository(db *sql.DB) *authSubjectRepository {
	return &authSubjectRepository{
		db: db,
	}
}

func (u *authSubjectRepository) Create(ctx context.Context, user *db.AuthSubject) error {
	return user.Insert(ctx, u.db, boil.Infer())
}
