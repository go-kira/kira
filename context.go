package kira

import (
	"fmt"
	"github.com/go-kira/kira/modules/log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-kira/kira/modules/config"
)

// Context pool
var contextPool = &sync.Pool{
	New: func() interface{} {
		return &Context{}
	},
}

// HandlerFunc - Type to define context function
type HandlerFunc func(*Context)

// Context ...
type Context struct {
	request  *http.Request
	response http.ResponseWriter
	logger   *log.Logger
	configs  *config.Config
	// The data associated with the request.
	data       map[string]interface{}
	statusCode int
	requestID  string
	startAt    time.Time
	// environment
	env string
}

// NewContext - Create new instance of Context
// func NewContext(w http.ResponseWriter, r *http.Request, app *App) *Context {
// 	ctx := contextPool.Get().(*Context)
// 	ctx.response = w
// 	ctx.request = r
// 	ctx.Logger = app.logger
// 	ctx.Configs = app.Configs
// 	ctx.data = make(map[string]interface{})
// 	ctx.env = app.Env
//
// 	return ctx
// }

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

// Redirect replies to the request with a redirect to url,
func (c *Context) Redirect(url string, code int) {
	http.Redirect(c.Response(), c.Request(), url, code)
}

// Log gets the Log instance.
func (c *Context) Log() *log.Logger {
	return setupLogger(c.Config(), c.logger.Writer, log.Fields{
		"status":     strconv.Itoa(c.StatusCode()),
		"method":     c.Request().Method,
		"path":       c.Request().RequestURI,
		"duration":   time.Since(c.startAt).String(),
		"request_id": c.RequestID(),
	})
}

// Config gets the application configs.
func (c *Context) Config() *config.Config {
	return c.configs
}

// Code sets response status statusCode.
func (c *Context) SetStatusCode(code int) {
	c.statusCode = code
}

// Code gets response status statusCode.
func (c *Context) StatusCode() int {
	return c.statusCode
}

// Code sets response status statusCode.
func (c *Context) SetRequestID(id string) {
	c.requestID = id
}

// Code gets response status statusCode.
func (c *Context) RequestID() string {
	return c.requestID
}

// Env gets the application environment.
func (c *Context) Env() string {
	return c.env
}

// Error stop the request with panic
func (c *Context) Error(msg interface{}) {
	// Just panic and the recover will come to save us :)
	// TODO: later we need something better than this.
	panic(fmt.Sprint(msg))
}

// Err checks if the error not empty.
// It's will redirect the error to Error method if there an error.
func (c *Context) Err(err error) {
	if err != nil {
		c.Error(err)
	}
}
