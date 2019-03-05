package kira

import (
	"fmt"
	"net/http"

	"github.com/go-kira/kog"
	"github.com/go-kira/kon"
)

// ContextFunc - Type to define context function
type ContextFunc func(*Context)

// Context ...
type Context struct {
	request  *http.Request
	response http.ResponseWriter
	Logger   *kog.Logger
	Configs  *kon.Kon
	// The data assocaited with the request.
	data map[string]interface{}
	// Will hold the response status code.
	statusCode int
}

// NewContext - Create new instance of Context
func NewContext(res http.ResponseWriter, req *http.Request, app *App) *Context {
	return &Context{
		request:  req,
		response: res,
		Logger:   app.Log,
		Configs:  app.Configs,
		data:     make(map[string]interface{}),
	}
}

// Request a Request represents an HTTP request received by a server.
func (c *Context) Request() *http.Request {
	return c.request
}

// Response is used by an HTTP handler to construct an HTTP response.
func (c *Context) Response() http.ResponseWriter {
	return c.response
}

// Header Write HTTP header to the response.
func (c *Context) Header(code int) {
	c.Response().WriteHeader(code)
}

// Redirect replies to the request with a redirect to url,
func (c *Context) Redirect(url string, code int) {
	http.Redirect(c.Response(), c.Request(), url, code)
}

// Log gets the Log instance.
func (c *Context) Log() *kog.Logger {
	return c.Logger
}

// Config gets the application configs.
func (c *Context) Config() *kon.Kon {
	return c.Configs
}

// Status send a specific status with the HTTP reply.
func (c *Context) Status(code int) {
	c.Response().WriteHeader(code)

	fmt.Fprint(c.Response(), http.StatusText(code))
}

// Error stop the request with panic
func (c *Context) Error(msg ...interface{}) {
	// Just panic and the recover will come to save us :)
	panic(fmt.Sprint(msg...))
}
