// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// TestGet implements GET /test operation.
//
// GET /test
func (UnimplementedHandler) TestGet(ctx context.Context) (r *TestGetOK, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrRespStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrRespStatusCode) {
	r = new(ErrRespStatusCode)
	return r
}
