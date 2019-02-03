package kira

import (
	"fmt"
	"net/http"

	"github.com/go-kira/kog"
)

// Example:
// app.Get("/", func (ctx *kira.Context) {
//
// })

// ContextFunc - Type to define context function
type ContextFunc func(*Context)

// Context ...
type Context struct {
	request  *http.Request
	response http.ResponseWriter
	Logger   *kog.Logger
	// The data assocaited with the request.
	data map[string]interface{}
	// Will hold the response status code.
	statusCode int
}

// NewContext - Create new instance of Context
func NewContext(res http.ResponseWriter, req *http.Request, logger *kog.Logger) *Context {
	return &Context{
		request:  req,
		response: res,
		Logger:   logger,
		data:     make(map[string]interface{}),
	}
}

// Request - get the request
func (c *Context) Request() *http.Request {
	return c.request
}

// Response - get the response
func (c *Context) Response() http.ResponseWriter {
	return c.response
}

// Header - Write http header to the response
func (c *Context) Header(code int) {
	c.Response().WriteHeader(code)
}

// Log - write the log
func (c *Context) Log() *kog.Logger {
	return c.Logger
}

// Status - Send specific status to the response.
func (c *Context) Status(code int) {
	c.Response().WriteHeader(code)

	fmt.Fprint(c.Response(), http.StatusText(code))
}
