package requestid

import (
	"context"

	"github.com/go-kira/kira"
	"github.com/google/uuid"
)

// RequestID struct
type RequestID struct{}

// New - new instance of RequestID.
func New() *RequestID {
	return &RequestID{}
}

// Middleware handler.
func (rq *RequestID) Middleware(ctx *kira.Context, next kira.HandlerFunc) {
	headerName := ctx.Config().GetString("server.request_id", "X-Request-Id")

	// Request ID
	requestid := rq.random()

	// Context
	requestIDContext := context.WithValue(ctx.Request().Context(), headerName, requestid)

	// Set header.
	ctx.Response().Header().Set(headerName, requestid)

	// Change the request with the new one with context.
	ctx.SetRequest(ctx.Request().WithContext(requestIDContext))

	// Move to the next handler.
	next(ctx)
}

// random return random string for request id
func (rq *RequestID) random() string {
	return uuid.New().String()
}
