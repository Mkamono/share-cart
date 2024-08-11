package handler

import (
	dbRepo "api/app/infra/repository"
	"api/app/oas"
	"api/app/usecase"
	"context"
	"log/slog"
)

func (h *handler) SignUpPost(ctx context.Context, req *oas.SignUpPostReq) (oas.SignUpPostRes, error) {
	sub := h.jwtClient.GetSubject(ctx)

	if sub == "" {
		return &oas.R401UnauthorizedError{
			Message: "Unauthorized",
		}, nil
	}

	txRepo := dbRepo.NewTransactionRepository(h.db)
	userRepo := dbRepo.NewUserRepository(h.db)
	authSubjectRepo := dbRepo.NewAuthSubjectRepository(h.db)

	createUserUsecase := usecase.NewCreateUserUsecase(txRepo, userRepo, authSubjectRepo)

	err := createUserUsecase.Run(ctx, req.Name.Value, sub)
	if err != nil {
		slog.ErrorContext(ctx, "Usecase: Failed to create user", "error", err)
		return &oas.R500InternalServerError{
			Message: "Internal server error",
		}, nil
	}

	slog.InfoContext(ctx, "Usecase: Create user success")
	return &oas.DefaultSuccess{
		Message: "Sign up success",
	}, nil
}
