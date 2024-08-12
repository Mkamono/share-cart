package jwt

import (
	"api/app/handler"
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type contextKey string

const subjectContextKey contextKey = "subject"

var _ handler.JwtClient = (*jwtClient)(nil)

type jwtClient struct {
	validator *validator.Validator
}

func NewClient(issuerURL *url.URL, audience []string) (*jwtClient, error) {
	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	v, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		audience,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create JWT validator: %v", err)
	}

	return &jwtClient{
		validator: v,
	}, nil
}

func (j *jwtClient) Validate(ctx context.Context, token string) (*validator.ValidatedClaims, error) {
	i, err := j.validator.ValidateToken(ctx, token)
	if err != nil {
		return nil, err
	}

	claims, ok := i.(*validator.ValidatedClaims)
	if !ok {
		return nil, fmt.Errorf("failed to assert validated claims type")
	}

	if len(claims.RegisteredClaims.Subject) == 0 {
		return nil, fmt.Errorf("subject in JWT claims is empty")
	}

	return claims, nil
}

func (j *jwtClient) WithSubjectContext(ctx context.Context, subject string) context.Context {
	return context.WithValue(ctx, subjectContextKey, subject)
}

func (j *jwtClient) GetSubjectFromContext(ctx context.Context) (string, bool) {
	subject, ok := ctx.Value(subjectContextKey).(string)
	return subject, ok
}
