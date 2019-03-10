package kira

// TODO:
//  - Remove "csrf", "session" from the base code.
//  - Implement "plugin" mechanism.
//  - We can use "plugin" to provide additional functionalities to the user like: Auth, Cache, Database ORM...
//  - Error wrapper: Error{op: "op.name", err: Error}...

import (
	"fmt"
	"net/http"

	"github.com/go-kira/config"
	"github.com/go-kira/log"
	"github.com/julienschmidt/httprouter"
)

var hero = `   __ __   _             
  / //_/  (_)  ____ ___ _
 / ,<    / /  / __// _  /
/_/|_|  /_/  /_/   \_,_/ 
`

// some bytes :)
const (
	KB = 1 << 10
	MB = 1 << 20
	GB = 1 << 30
)

// App hold the framework options
type App struct {
	Routes      []*Route
	Middlewares []Middleware
	Router      *httprouter.Router
	Log         *log.Logger
	Configs     *config.Config
	Env         string

	// Not found handler.
	NotFoundHandler HandlerFunc
}

// New init the framework
func New() *App {
	// initialization...
	app := &App{}

	// kira environment
	app.Env = getEnv()

	// configs
	app.Configs = getConfig()

	// logger
	app.Log = setupLogger(app.Configs)

	// define a Router
	app.Router = httprouter.New()

	// return App instance
	return app
}

// Run the framework
func (a *App) Run() *App {
	fmt.Printf("%v", hero)

	a.RegisterRoutes()

	// validate if the server need tls connection.
	if !a.Configs.GetBool("server.tls", false) {
		// Start the server
		a.StartServer()
	} else {
		// TLS
		a.StartTLSServer()
	}

	// return App instance
	return a
}

// NotFound custom not found handler.
func (a *App) NotFound(ctx HandlerFunc) {
	a.NotFoundHandler = ctx
}

// default not found handler.
func defaultNotFound(ctx *Context) {
	ctx.HeaderStatus(http.StatusNotFound)

	// JSON
	if ctx.WantsJSON() {
		// Json response
		ctx.JSON(struct {
			Error   int    `json:"error"`
			Message string `json:"message"`
		}{http.StatusNotFound, "404 Not Found"})
		return
	}

	// HTML
	// Validate if the template exists
	if ctx.ViewExists("errors/404") {
		ctx.View("errors/404")
	} else {
		ctx.String("<!DOCTYPE html><html><head><title>404 Not Found</title></head><body>404 Not Found</body></html>")
	}
}
