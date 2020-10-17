package kira

import (
	"github.com/google/uuid"
)

// RequestID struct
type RequestID struct{}

// new instance of RequestID.
func NewRequestID() *RequestID {
	return &RequestID{}
}

// Middleware handler.
func (rq *RequestID) Middleware(ctx *Context, next HandlerFunc) {
	headerName := ctx.Config().GetString("server.request_id", "X-Request-Id")

	// Set header.
	ctx.Response().Header().Set(headerName, ctx.RequestID())

	// Move to the next handler.
	next(ctx)
}

// random return random string for request requestID
func (rq *RequestID) random() string {
	return uuid.New().String()
}
