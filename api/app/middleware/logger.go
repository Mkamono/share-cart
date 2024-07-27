package middleware

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ogen-go/ogen/middleware"
)

func NewLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func Logging(logger *slog.Logger) middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {
		logger.Info(
			"Handling request",
			slog.String("operation", req.OperationName),
			slog.String("URL", req.Raw.URL.String()),
			slog.String("operationId", req.OperationID),
			slog.String("body", fmt.Sprintf("%#v", req.Body)),
			slog.String("raw request", fmt.Sprintf("%+v", req.Raw)),
		)
		resp, err := next(req)
		if err != nil {
			logger.Error("Fail", "error", err)
		} else {
			logger.Info(
				"Handled request",
				"operation", req.OperationName,
				"operationId", req.OperationID,
				"raw response", fmt.Sprintf("%#v", resp.Type),
			)
		}
		return resp, err
	}
}
