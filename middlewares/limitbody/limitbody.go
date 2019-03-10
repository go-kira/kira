package limitbody

import (
	"net/http"

	"github.com/go-kira/kira"
)

// MB - one MB.
const MB = 1 << 20

// Limitbody - Middleware.
type Limitbody struct{}

// New - return Limitbody instance
func New() *Limitbody {
	return &Limitbody{}
}

// Middleware handler.
func (l *Limitbody) Middleware(ctx *kira.Context, next kira.HandlerFunc) {
	if ctx.Request().ContentLength > ctx.Config().GetInt64("server.body_limit", 32)*MB {
		http.Error(ctx.Response(), "Request too large", http.StatusExpectationFailed)
		return
	}
	ctx.Request().Body = http.MaxBytesReader(ctx.Response(), ctx.Request().Body, ctx.Config().GetInt64("server.body_limit", 32)*MB)

	next(ctx)
}
