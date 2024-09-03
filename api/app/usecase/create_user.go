package usecase

import (
	dbEntity "api/app/domain/entity/db"
	dbRepo "api/app/domain/repository/db"
	"context"
	"log/slog"
)

type CreateUserUsecase interface {
	Run(ctx context.Context, sub string, name string) error
}

var _ CreateUserUsecase = (*createUserUsecase)(nil)

type createUserUsecase struct {
	txRepo      dbRepo.TransactionRepository
	userRepo    dbRepo.UserRepository
	subjectRepo dbRepo.AuthSubjectRepository
}

func NewCreateUserUsecase(
	txRepo dbRepo.TransactionRepository,
	userRepo dbRepo.UserRepository,
	subjectRepo dbRepo.AuthSubjectRepository,
) CreateUserUsecase {
	return &createUserUsecase{
		txRepo:      txRepo,
		userRepo:    userRepo,
		subjectRepo: subjectRepo,
	}
}

func (u *createUserUsecase) Run(ctx context.Context, name string, sub string) error {
	err := u.txRepo.WithInTx(ctx, func(ctx context.Context) error {
		user := &dbEntity.User{}
		user.Name = name
		err := u.userRepo.Create(ctx, user)
		if err != nil {
			slog.ErrorContext(ctx, "Repo: Failed to create user", "error", err)
			return err
		}
		slog.InfoContext(ctx, "Repo: Success to create user", "user", struct {
			ID   string
			Name string
		}{
			ID:   user.ID,
			Name: user.Name,
		})

		subject := &dbEntity.AuthSubject{}
		subject.Subject = sub
		subject.UserID = user.ID
		err = u.subjectRepo.Create(ctx, subject)
		if err != nil {
			slog.ErrorContext(ctx, "Repo: Failed to create auth subject", "error", err)
			return err
		}
		slog.InfoContext(ctx, "Repo: Success to create auth subject")

		return nil
	})
	if err != nil {
		slog.ErrorContext(ctx, "Repo: Failed to run transaction", "error", err)
		return err
	}
	slog.InfoContext(ctx, "Repo: Success to run transaction")

	return nil
}
