package jwt

import (
	"api/app/handler"
	"context"
)

type jwtClient struct {
	subjectKey string
}

func NewJwtClient(subjectKey string) handler.JwtClient {
	return &jwtClient{subjectKey: subjectKey}
}

func (j *jwtClient) GetSubject(ctx context.Context) string {
	return ctx.Value(j.subjectKey).(string)
}
