package middleware

import (
	"context"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger(logglerFunc func(c echo.Context, v middleware.RequestLoggerValues) error) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:     true,
		LogURI:        true,
		LogError:      true,
		HandleError:   true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: logglerFunc,
	})
}

func SlogLoggerFunc(c echo.Context, v middleware.RequestLoggerValues) error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if v.Error == nil {
		logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
			slog.String("uri", v.URI),
			slog.Int("status", v.Status),
		)
	} else {
		logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
			slog.String("uri", v.URI),
			slog.Int("status", v.Status),
			slog.String("err", v.Error.Error()),
		)
	}
	return nil
}
