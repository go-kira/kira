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

func (e *Example) Name() string {
	return "example"
}

// Middleware ...
func (e *Example) Middleware(c *kira.Context, next kira.HandlerFunc) {
	// Next handlerr
	c.String("before \n")

	next(c)

	c.String("after \n")
	// next.ServeHTTP(ctx.Response(), ctx.Request())
}
