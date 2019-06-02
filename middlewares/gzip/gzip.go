package gzip

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/go-kira/kira"
)

var gzPool sync.Pool

// New ...
func New() Gzip {
	return Gzip{}
}

// Gzip ...
type Gzip struct{}

// Middleware ...
func (g Gzip) Middleware(ctx *kira.Context, next kira.HandlerFunc) {
	gzPool.New = func() interface{} {
		gz, err := gzip.NewWriterLevel(nil, ctx.Config().GetInt("gzip.level", gzip.DefaultCompression))
		if err != nil {
			panic(err)
		}
		return gz
	}

	if !strings.Contains(ctx.Request().Header.Get("Accept-Encoding"), "gzip") {
		next(ctx)
		return
	}

	// GZip
	gz := gzPool.Get().(*gzip.Writer)
	defer gzPool.Put(gz)
	defer gz.Reset(nil)
	gz.Reset(ctx.Response())

	ctx.Response().Header().Set("Content-Encoding", "gzip")
	ctx.Response().Header().Set("Vary", "Accept-Encoding")
	ctx.SetResponse(&gzipResponseWriter{
		Writer:         gz,
		ResponseWriter: ctx.Response(),
	})
	defer func() {
		gz.Close()
	}()

	// Next to the next handler.
	next(ctx)
}

// Custom ResponseWriter for gzip
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *gzipResponseWriter) WriteHeader(code int) {
	w.Header().Del("Content-Length")
	w.ResponseWriter.WriteHeader(code)
}
func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	log.Println("gzip: type:" + w.Header().Get("Content-Type"))
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return w.Writer.Write(b)
}
