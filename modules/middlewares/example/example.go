package example

import (
	"github.com/go-kira/kira"
)

// Example - kira middleware example.
type Example struct{}

// New - a new instance of Example
func New() *Example {
	return &Example{}
}

// Middleware handler.
func (e *Example) Middleware(c *kira.Context, next kira.HandlerFunc) {
	// Next handlerr
	c.WriteString("before \n")

	next(c)

	c.WriteString("after \n")
	// next.ServeHTTP(ctx.Response(), ctx.Request())
}
