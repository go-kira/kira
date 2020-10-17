package kira

// Log - log middleware
type Log struct{}

// New ...
func NewLogger() *Log {
	return &Log{}
}

// Middleware handler.
func (l *Log) Middleware(ctx *Context, next HandlerFunc) {
	next(ctx)

	ctx.Log().Info("Request end")
}
