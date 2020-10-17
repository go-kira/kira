package kira

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Log - log middleware
type Log struct{}

// New ...
func NewLogger() *Log {
	return &Log{}
}

// Middleware handler.
func (l *Log) Middleware(ctx *Context, next HandlerFunc) {
	// Start time
	var now = time.Now()
	// Store the status code
	statusRecorder := &statusRecorder{ctx.Response(), http.StatusOK}

	// Change the ResponseWriter to our recorder.
	ctx.SetResponse(statusRecorder)

	next(ctx)

	logFormat := ctx.Config().GetString("log.format", ":status :method :duration :request_id :path")

	r := strings.NewReplacer(
		// ":time", log.FormatTime(time.Now()),
		":status", strconv.Itoa(statusRecorder.statusCode),
		":method", ctx.Request().Method,
		":path", ctx.Request().RequestURI,
		":duration", time.Since(now).String(),
		":request_id", ctx.Request().Context().Value(ctx.Config().GetString("server.request_id", "X-Request-Id")).(string),
	)
	ctx.Log().Info(r.Replace(logFormat))
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader - store the header to use it later.
func (r *statusRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}
