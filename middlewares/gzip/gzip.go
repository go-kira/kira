package gzip

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/go-kira/kira"
)

// New ...
func New() Gzip {
	return Gzip{}
}

// Gzip ...
type Gzip struct{}

// Middleware ...
func (g Gzip) Middleware(ctx *kira.Context, next kira.HandlerFunc) {
	ctx.Response().Header().Add("Vary", "Accept-Encoding")
	if strings.Contains(ctx.Request().Header.Get("Accept-Encoding"), "gzip") {
		ctx.Response().Header().Set("Content-Encoding", "gzip")
		gz, err := gzip.NewWriterLevel(ctx.Response(), ctx.Config().GetInt("gzip.level", gzip.DefaultCompression))
		if err != nil {
			ctx.Error(err)
		}
		defer gz.Close()

		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: ctx.Response()}
		ctx.SetResponse(gzr)

		next(ctx)
	} else {
		next(ctx)
		return
	}
}

// Custom ResponseWriter for gzip
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) WriteHeader(code int) {
	if code == http.StatusNoContent {
		w.ResponseWriter.Header().Del("Content-Encoding")
	}
	w.Header().Del("Content-Length")
	w.ResponseWriter.WriteHeader(code)
}
func (w gzipResponseWriter) Write(b []byte) (int, error) {
	// if w.Header().Get("Content-Type") == "" {
	// 	w.Header().Set("Content-Type", http.DetectContentType(b))
	// }
	return w.Writer.Write(b)
}
func (w gzipResponseWriter) Flush() {
	w.Writer.(*gzip.Writer).Flush()
}
