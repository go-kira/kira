package gzip

import (
	"bufio"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net"
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
		gz, err := gzip.NewWriterLevel(ioutil.Discard, ctx.Config().GetInt("gzip.level", gzip.DefaultCompression))
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
	defer gz.Reset(ioutil.Discard)
	gz.Reset(ctx.Response())

	ctx.Response().Header().Set("Content-Encoding", "gzip")
	ctx.Response().Header().Set("Vary", "Accept-Encoding")
	ctx.SetResponse(&gzipWriter{
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
type gzipWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *gzipWriter) WriteHeader(code int) {
	if code == http.StatusNoContent {
		w.ResponseWriter.Header().Del("Content-Encoding")
	}
	w.Header().Del("Content-Length")
	w.ResponseWriter.WriteHeader(code)
}
func (w *gzipWriter) Write(b []byte) (int, error) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}

	return w.Writer.Write(b)
}
func (w *gzipWriter) Flush() {
	w.Writer.(*gzip.Writer).Flush()
	if flusher, ok := w.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	}
}
func (w *gzipWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}
