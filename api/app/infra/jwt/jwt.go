package jwt

import (
	"api/app/handler"
	"context"
	"log/slog"
	"os"
)

type jwtClient struct {
	subjectKey string
}

func NewJwtClient(subjectKey string) handler.JwtClient {
	return &jwtClient{subjectKey: subjectKey}
}

func (j *jwtClient) GetSubject(ctx context.Context) string {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	v := ctx.Value(j.subjectKey)

	subject, ok := v.(string)
	if !ok {
		logger.Error("failed to get subject from context")
		return ""
	}

	return subject
}
