package middleware

import (
	"api/app/shared"
	"api/app/shared/ctxlogger"
	"context"
	"fmt"
	"log/slog"
	"regexp"

	"github.com/ogen-go/ogen/middleware"
)

const (
	googleCloudTraceHeader = "X-Cloud-Trace-Context"
)

func NewRequestLogger(projectID string) middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {
		req.Context = ctxlogger.WithValue(req.Context, "operation", req.OperationName)
		traceHeader := req.Raw.Header.Get(googleCloudTraceHeader)
		req.Context = addTraceID(req.Context, traceHeader, projectID)

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
			slog.InfoContext(req.Context, "Returned response")
		}

		return resp, err
	}
}

func addTraceID(ctx context.Context, traceHeader string, projectID string) context.Context {
	traceID, ok := extractTraceID(traceHeader)
	if ok {
		trace := fmt.Sprintf("projects/%s/traces/%s", projectID, traceID)
		return ctxlogger.WithValue(ctx, "logging.googleapis.com/trace", trace)
	} else {
		fmt.Println("traceID", "not found")
		trace := fmt.Sprintf("projects/%s/traces/%s", projectID, shared.Uuid())
		return ctxlogger.WithValue(ctx, "logging.googleapis.com/trace", trace)
	}
}

func extractTraceID(header string) (string, bool) {
	if header == "" {
		return "", false
	}

	// NOTE: https://cloud.google.com/trace/docs/setup?hl=ja#force-trace
	matches := regexp.MustCompile(`([a-f\d]+)/([a-f\d]+)`).FindAllSubmatch([]byte(header), -1)
	if len(matches) != 1 {
		return "", false
	}

	sub := matches[0]
	if len(sub) != 3 {
		return "", false
	}

	return string(sub[1]), true
}
