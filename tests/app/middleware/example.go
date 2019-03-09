package middleware

import "github.com/go-kira/kira"

// Example
type Example struct{}

func New() *Example {
	return &Example{}
}
func (e *Example) Middleware(ctx *kira.Context, next kira.HandlerFunc) {
	ctx.String("Before")
	next(ctx)
	ctx.String("After")
}

// Example2
type Example2 struct{}

func New2() *Example2 {
	return &Example2{}
}
func (e *Example2) Middleware(ctx *kira.Context, next kira.HandlerFunc) {
	ctx.String("Before2")
	next(ctx)
	ctx.String("After2")
}

// ContextData
type ContextData struct{}

func NewContextData() *ContextData {
	return &ContextData{}
}
func (e *ContextData) Middleware(ctx *kira.Context, next kira.HandlerFunc) {
	// We should see this inside a normal handler that use this middleware.
	ctx.SetData("foo", "bar")

	next(ctx)
}
