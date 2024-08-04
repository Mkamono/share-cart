package handler

import (
	"api/app/oas"
	"context"
	"fmt"
	"time"
)

var _ oas.Handler = (*handler)(nil)

func NewHandler(subjectKey string) oas.Handler {
	return &handler{
		getSubject: func(ctx context.Context) string {
			return ctx.Value(subjectKey).(string)
		},
	}
}

type handler struct {
	// getSubject is a function that returns the subject from the context.
	getSubject func(context.Context) string
}

func (h *handler) TestGet(context.Context) (oas.TestGetRes, error) {
	s := fmt.Sprintf("Hello, world! : %s", time.Now().String())
	return &oas.TestGetOK{
		Message: oas.OptString{Value: s, Set: true},
	}, nil
}

func (h *handler) LoginPost(ctx context.Context, req *oas.LoginPostReq) (oas.LoginPostRes, error) {
	sub := h.getSubject(ctx)
	return &oas.LoginPostOK{
		Message: oas.OptString{Value: sub, Set: true},
	}, nil
}
