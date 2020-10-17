package kira

import "net/http"

// responseWriter for kira framework.
type responseWriter struct {
	http.ResponseWriter
	ctx *Context
}

// WriteHeader - store the header to use it later.
func (recorder *responseWriter) WriteHeader(code int) {
	recorder.ctx.SetStatusCode(code)
	recorder.ResponseWriter.WriteHeader(code)
}
