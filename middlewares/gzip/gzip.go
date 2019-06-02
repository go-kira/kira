package gzip

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

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
	var gzPool sync.Pool
	gzPool.New = func() interface{} {
		gz, err := gzip.NewWriterLevel(ioutil.Discard, ctx.Config().GetInt("gzip.level", gzip.DefaultCompression))
		if err != nil {
			panic(err)
		}
		return gz
	}

	if strings.Contains(ctx.Request().Header.Get("Accept-Encoding"), "gzip") {
		log.Println("gzip: compressing")
		gz := gzPool.Get().(*gzip.Writer)
		defer gzPool.Put(gz)
		defer gz.Reset(ioutil.Discard)
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
	} else {
		log.Println("gzip: no compress")
	}
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
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return w.Writer.Write(b)
}
