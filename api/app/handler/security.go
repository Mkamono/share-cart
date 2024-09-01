package handler

import (
	"api/app/oas"
	"context"
)

var _ oas.SecurityHandler = (*securityHandler)(nil)

type securityHandler struct {
	jwtClient JwtClient
}

func (s *securityHandler) HandleBearer(ctx context.Context, operationName string, t oas.Bearer) (context.Context, error) {
	claims, err := s.jwtClient.Validate(ctx, t.Token)
	if err != nil {
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
