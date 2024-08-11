package middleware

import (
	"api/app/shared/ctxlogger"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/ogen-go/ogen/middleware"
)

func NewSlogLogger() middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {
		traceID := uuid.New()
		req.Context = ctxlogger.WithValue(req.Context, "traceID", traceID.String())
		req.Context = ctxlogger.WithValue(req.Context, "operation", req.OperationName)

		slog.InfoContext(req.Context, "Received request", "request", struct {
			URL  string
			Body string
		}{
			URL:  req.Raw.URL.String(),
			Body: fmt.Sprintf("%#v", req.Body),
		})

		resp, err := next(req)
		if err != nil {
			slog.ErrorContext(req.Context, "Error", "error", err)
		} else {
			slog.InfoContext(req.Context, "Returned response", "response", resp.Type)
		}

		return resp, err
	}
}
