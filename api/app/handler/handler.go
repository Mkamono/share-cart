package handler

import (
	"api/app/oas"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"
)

var _ oas.Handler = (*handler)(nil)

type JwtClient interface {
	GetSubject(ctx context.Context) string
}

func NewHandler(
	jc JwtClient,
	l *slog.Logger,
	db *sql.DB,
) oas.Handler {
	return &handler{
		jwtClient: jc,
		logger:    l,
		db:        db,
	}
}

type handler struct {
	jwtClient JwtClient
	logger    *slog.Logger
	db        *sql.DB
}

func (h *handler) TestGet(ctx context.Context) (oas.TestGetRes, error) {
	s := fmt.Sprintf("Hello, world! : %s", time.Now().String())
	return &oas.DefaultSuccess{
		Message: s,
	}, nil
}
