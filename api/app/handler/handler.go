package handler

import (
	"api/app/oas"
	"context"
	"database/sql"
	"fmt"
	"time"
)

var _ oas.Handler = (*handler)(nil)

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

func (h *handler) NewError(ctx context.Context, err error) *oas.R500InternalServerErrorStatusCode {
	return &oas.R500InternalServerErrorStatusCode{
		StatusCode: 500,
		Response: oas.R500InternalServerError{
			Message: err.Error(),
		},
	}
}

func (h *handler) TestGet(ctx context.Context) (*oas.R200OK, error) {
	s := fmt.Sprintf("Hello, world! : %s", time.Now().String())
	return &oas.R200OK{
		Message: s,
	}, nil
}
