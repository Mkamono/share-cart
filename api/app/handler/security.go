package handler

import (
	"api/app/oas"
	"context"
	"log/slog"

	"github.com/auth0/go-jwt-middleware/v2/validator"
)

var _ oas.SecurityHandler = (*securityHandler)(nil)

type JwtClient interface {
	Validate(ctx context.Context, token string) (*validator.ValidatedClaims, error)
	WithSubjectContext(ctx context.Context, subject string) context.Context
	GetSubjectFromContext(ctx context.Context) (string, bool)
}

type securityHandler struct {
	jwtClient JwtClient
}

func (s *securityHandler) HandleBearer(ctx context.Context, operationName string, t oas.Bearer) (context.Context, error) {
	claims, err := s.jwtClient.Validate(ctx, t.Token)
	if err != nil {
		slog.WarnContext(ctx, "Security: Failed to validate token", "error", err)
		return ctx, nil
	}

	ctx = s.jwtClient.WithSubjectContext(ctx, claims.RegisteredClaims.Subject)
	return ctx, nil
}

func NewSecurityHandler(jwtClient JwtClient) oas.SecurityHandler {
	return &securityHandler{
		jwtClient: jwtClient,
	}
}
