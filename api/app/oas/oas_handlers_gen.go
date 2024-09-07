// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
)

// handleMarketGetRequest handles GET /market operation.
//
// Get all markets.
//
// GET /market
func (s *Server) handleMarketGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/market"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MarketGet",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "MarketGet",
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearer(ctx, "MarketGet", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "Bearer",
					Err:              err,
				}
				if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
					defer recordError("Security:Bearer", err)
				}
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
				defer recordError("Security", err)
			}
			return
		}
	}

	var response []Market
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "MarketGet",
			OperationSummary: "Get all markets",
			OperationID:      "",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = []Market
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.MarketGet(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.MarketGet(ctx)
	}
	if err != nil {
		if errRes, ok := errors.Into[*R500InternalServerErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w, span); err != nil {
				defer recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w, span); err != nil {
			defer recordError("Internal", err)
		}
		return
	}

	if err := encodeMarketGetResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleMarketMarketIdDeleteRequest handles DELETE /market/{marketId} operation.
//
// Delete a market by ID.
//
// DELETE /market/{marketId}
func (s *Server) handleMarketMarketIdDeleteRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/market/{marketId}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MarketMarketIdDelete",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "MarketMarketIdDelete",
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearer(ctx, "MarketMarketIdDelete", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "Bearer",
					Err:              err,
				}
				if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
					defer recordError("Security:Bearer", err)
				}
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
				defer recordError("Security", err)
			}
			return
		}
	}
	params, err := decodeMarketMarketIdDeleteParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *MarketMarketIdDeleteNoContent
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "MarketMarketIdDelete",
			OperationSummary: "Delete a market by ID",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "marketId",
					In:   "path",
				}: params.MarketId,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = MarketMarketIdDeleteParams
			Response = *MarketMarketIdDeleteNoContent
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackMarketMarketIdDeleteParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.MarketMarketIdDelete(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.MarketMarketIdDelete(ctx, params)
	}
	if err != nil {
		if errRes, ok := errors.Into[*R500InternalServerErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w, span); err != nil {
				defer recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w, span); err != nil {
			defer recordError("Internal", err)
		}
		return
	}

	if err := encodeMarketMarketIdDeleteResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleMarketMarketIdGetRequest handles GET /market/{marketId} operation.
//
// Get a market by ID.
//
// GET /market/{marketId}
func (s *Server) handleMarketMarketIdGetRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/market/{marketId}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MarketMarketIdGet",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "MarketMarketIdGet",
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearer(ctx, "MarketMarketIdGet", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "Bearer",
					Err:              err,
				}
				if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
					defer recordError("Security:Bearer", err)
				}
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
				defer recordError("Security", err)
			}
			return
		}
	}
	params, err := decodeMarketMarketIdGetParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *Market
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "MarketMarketIdGet",
			OperationSummary: "Get a market by ID",
			OperationID:      "",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "marketId",
					In:   "path",
				}: params.MarketId,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = MarketMarketIdGetParams
			Response = *Market
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackMarketMarketIdGetParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.MarketMarketIdGet(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.MarketMarketIdGet(ctx, params)
	}
	if err != nil {
		if errRes, ok := errors.Into[*R500InternalServerErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w, span); err != nil {
				defer recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w, span); err != nil {
			defer recordError("Internal", err)
		}
		return
	}

	if err := encodeMarketMarketIdGetResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleMarketPostRequest handles POST /market operation.
//
// Create a new market.
//
// POST /market
func (s *Server) handleMarketPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/market"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "MarketPost",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "MarketPost",
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearer(ctx, "MarketPost", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "Bearer",
					Err:              err,
				}
				if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
					defer recordError("Security:Bearer", err)
				}
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
				defer recordError("Security", err)
			}
			return
		}
	}
	request, close, err := s.decodeMarketPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response *Market
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "MarketPost",
			OperationSummary: "Create a new market",
			OperationID:      "",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *MarketPostReq
			Params   = struct{}
			Response = *Market
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.MarketPost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.MarketPost(ctx, request)
	}
	if err != nil {
		if errRes, ok := errors.Into[*R500InternalServerErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w, span); err != nil {
				defer recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w, span); err != nil {
			defer recordError("Internal", err)
		}
		return
	}

	if err := encodeMarketPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleSignUpPostRequest handles POST /sign-up operation.
//
// POST /sign-up
func (s *Server) handleSignUpPostRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/sign-up"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "SignUpPost",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "SignUpPost",
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearer(ctx, "SignUpPost", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "Bearer",
					Err:              err,
				}
				if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
					defer recordError("Security:Bearer", err)
				}
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
				defer recordError("Security", err)
			}
			return
		}
	}
	request, close, err := s.decodeSignUpPostRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response SignUpPostRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "SignUpPost",
			OperationSummary: "",
			OperationID:      "",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = *SignUpPostReq
			Params   = struct{}
			Response = SignUpPostRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.SignUpPost(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.SignUpPost(ctx, request)
	}
	if err != nil {
		if errRes, ok := errors.Into[*R500InternalServerErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w, span); err != nil {
				defer recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w, span); err != nil {
			defer recordError("Internal", err)
		}
		return
	}

	if err := encodeSignUpPostResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleTestGetRequest handles GET /test operation.
//
// GET /test
func (s *Server) handleTestGetRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		semconv.HTTPMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/test"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "TestGet",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "TestGet",
			ID:   "",
		}
	)
	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			sctx, ok, err := s.securityBearer(ctx, "TestGet", r)
			if err != nil {
				err = &ogenerrors.SecurityError{
					OperationContext: opErrContext,
					Security:         "Bearer",
					Err:              err,
				}
				if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
					defer recordError("Security:Bearer", err)
				}
				return
			}
			if ok {
				satisfied[0] |= 1 << 0
				ctx = sctx
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			err = &ogenerrors.SecurityError{
				OperationContext: opErrContext,
				Err:              ogenerrors.ErrSecurityRequirementIsNotSatisfied,
			}
			if encodeErr := encodeErrorResponse(s.h.NewError(ctx, err), w, span); encodeErr != nil {
				defer recordError("Security", err)
			}
			return
		}
	}

	var response *R200OK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "TestGet",
			OperationSummary: "",
			OperationID:      "",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *R200OK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.TestGet(ctx)
				return response, err
			},
		)
	} else {
		response, err = s.h.TestGet(ctx)
	}
	if err != nil {
		if errRes, ok := errors.Into[*R500InternalServerErrorStatusCode](err); ok {
			if err := encodeErrorResponse(errRes, w, span); err != nil {
				defer recordError("Internal", err)
			}
			return
		}
		if errors.Is(err, ht.ErrNotImplemented) {
			s.cfg.ErrorHandler(ctx, w, r, err)
			return
		}
		if err := encodeErrorResponse(s.h.NewError(ctx, err), w, span); err != nil {
			defer recordError("Internal", err)
		}
		return
	}

	if err := encodeTestGetResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
