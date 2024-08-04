package middleware

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/ogen-go/ogen/middleware"
	"golang.org/x/net/context"
)

// NewJwtMiddleware : JWTを検証するミドルウェアを作成します。contextにsubjectを追加します。検証に失敗した場合は空文字列を追加します。
func NewJwtMiddleware(auth0Domain string, auth0Audience string, subjectKey string) middleware.Middleware {
	issuer := fmt.Sprintf("https://%s/", auth0Domain)
	audience := []string{auth0Audience}

	jwtValidator, err := NewJwtValidator(issuer, audience)
	if err != nil {
		fmt.Println("failed to create jwt validator")
		panic(err)
	}

	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {
		tokenString := req.Raw.Header.Get("authorization")
		claims, err := jwtValidator.ValidateToken(req.Context, tokenString)

		var subject string
		if err != nil {
			subject = ""
		} else {
			subject = claims.RegisteredClaims.Subject
		}

		req.Context = context.WithValue(req.Context, subjectKey, subject)

		resp, err := next(req)
		return resp, err
	}
}

// JwtValidator is an interface for validating JWT tokens.
// 可読性のために、このインターフェースを定義しています。
type JwtValidator interface {
	ValidateToken(ctx context.Context, token string) (*validator.ValidatedClaims, error)
}

var _ JwtValidator = (*jwtValidator)(nil)

type jwtValidator struct {
	validator *validator.Validator
}

func NewJwtValidator(issuer string, audience []string) (JwtValidator, error) {
	issuerURL, err := url.Parse(issuer)
	if err != nil {
		log.Fatalf("failed to parse the issuer url: %v", err)
	}
	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	v, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		audience,
	)
	if err != nil {
		log.Fatalf("failed to set up the validator: %v", err)
	}

	return &jwtValidator{
		validator: v,
	}, nil
}

func (j *jwtValidator) ValidateToken(ctx context.Context, token string) (*validator.ValidatedClaims, error) {
	i, err := j.validator.ValidateToken(ctx, token)
	if err != nil {
		return nil, err
	}

	claims, ok := i.(*validator.ValidatedClaims)
	if !ok {
		return nil, fmt.Errorf("failed to get validated claims")
	}

	if len(claims.RegisteredClaims.Subject) == 0 {
		return nil, fmt.Errorf("subject in JWT claims was empty")
	}

	return claims, nil
}
