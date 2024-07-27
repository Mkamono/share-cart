package handler

import (
	"api/app/oas"
	"context"
	"fmt"
	"time"
)

var _ oas.Handler = (*handler)(nil)

func NewHandler() oas.Handler {
	return &handler{}
}

type handler struct {
}

func (h *handler) TestGet(context.Context) (*oas.TestGetOK, error) {
	s := fmt.Sprintf("Hello, world! : %s", time.Now().String())
	return &oas.TestGetOK{
		Message: oas.OptString{Value: s, Set: true},
	}, nil
}

func (h *handler) NewError(context.Context, error) *oas.ErrRespStatusCode {
	return &oas.ErrRespStatusCode{}
}
