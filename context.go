package kira

import (
	"fmt"
	"net/http"

	"github.com/go-kira/log"

	"github.com/go-kira/config"
)

// HandlerFunc - Type to define context function
type HandlerFunc func(*Context)

// Context ...
type Context struct {
	request  *http.Request
	response http.ResponseWriter
	Logger   *log.Logger
	Configs  *config.Config
	// The data associated with the request.
	data map[string]interface{}
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

// SetRequest change the current request with the given one.
func (c *Context) SetRequest(r *http.Request) {
	c.request = r
}

// SetResponse change the current response with the given one.
func (c *Context) SetResponse(w http.ResponseWriter) {
	c.response = w
}

// Request a Request represents an HTTP request received by a server.
func (c *Context) Request() *http.Request {
	return c.request
}

// Response is used by an HTTP handler to construct an HTTP response.
func (c *Context) Response() http.ResponseWriter {
	return c.response
}

// HeaderStatus Write HTTP header to the response and also write the status message to the body.
func (c *Context) HeaderStatus(code int) {
	c.Response().WriteHeader(code)
}

// Redirect replies to the request with a redirect to url,
func (c *Context) Redirect(url string, code int) {
	http.Redirect(c.Response(), c.Request(), url, code)
}

// Log gets the Log instance.
func (c *Context) Log() *log.Logger {
	return c.Logger
}

// Config gets the application configs.
func (c *Context) Config() *config.Config {
	return c.Configs
}

// Env gets the application environment.
func (c *Context) Env() string {
	return c.Env
}

// Status send a specific status with the HTTP reply.
func (c *Context) Status(code int) {
	c.Response().WriteHeader(code)

	fmt.Fprint(c.Response(), http.StatusText(code))
}

// Error stop the request with panic
func (c *Context) Error(msg interface{}, status ...int) {
	if len(status) > 0 {
		c.HeaderStatus(status[0])
	}

	// Just panic and the recover will come to save us :)
	panic(fmt.Sprint(msg))
}
