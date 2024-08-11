package handler

import (
	"api/app/oas"
	"context"
	"database/sql"
	"fmt"
	"time"
)

var _ oas.Handler = (*handler)(nil)

type JwtClient interface {
	GetSubject(ctx context.Context) string
}

func NewHandler(
	jc JwtClient,
	db *sql.DB,
) oas.Handler {
	return &handler{
		jwtClient: jc,
		db:        db,
	}
}

type handler struct {
	jwtClient JwtClient
	db        *sql.DB
}

func (h *handler) TestGet(ctx context.Context) (oas.TestGetRes, error) {
	s := fmt.Sprintf("Hello, world! : %s", time.Now().String())
	return &oas.DefaultSuccess{
		Message: s,
	}, nil
}
