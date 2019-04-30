package kira

// TODO:
//  - Implement "plugin" mechanism.
//  - We can use "plugin" to provide additional functionalities to the user like: Auth, Cache, Database ORM...
//  - Error wrapper: Error{op: "op.name", err: Error}...

import (
	"fmt"
	"net/http"
	"os"
	"sync"

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

// Map a type to represent map, this will be used alot in the internal code.
type Map map[string]interface{}

// App hold the framework options
type App struct {
	Routes      []*Route
	Middlewares []Middleware
	Router      *httprouter.Router
	Log         *log.Logger
	Configs     *config.Config
	Env         string

	// Not found handler
	NotFoundHandler HandlerFunc

	// Context pool
	pool *sync.Pool
}

// New init the framework
func New() *App {
	app := &App{}
	app.Env = getEnv()
	app.Configs = getConfig()
	app.Log = setupLogger(app.Configs)
	app.Router = httprouter.New()

	// Context pool
	app.pool = &sync.Pool{
		New: func() interface{} {
			return &Context{
				logger:  app.Log,
				configs: app.Configs,
				data:    make(map[string]interface{}),
				env:     app.Env,
			}
		},
	}

	// return App instance
	return app
}

// Run the framework
func (app *App) Run(config interface{}) *App {
	fmt.Printf("%v", hero)

	// Register the application routes
	app.RegisterRoutes()

	// Timezone
	tz := app.Configs.GetString("app.timezone")
	if tz != "" {
		os.Setenv("TZ", tz)
	}

	// Server
	var server *http.Server

	switch config.(type) {
	case *http.Server:
		server = config.(*http.Server)
		server.Handler = app.Router
	case string:
		server = &http.Server{
			Addr:    serverAddr(app.Configs, config.(string)),
			Handler: app.Router,
		}
	default:
		log.Panic("kira: Unspported type in Run args")
	}

	if !app.Configs.GetBool("server.tls", false) {
		app.StartServer(server)
	} else {
		app.StartTLSServer(server)
	}

	// App instance
	return app
}

// NotFound custom not found handler.
func (app *App) NotFound(ctx HandlerFunc) {
	app.NotFoundHandler = ctx
}

// default not found handler.
func defaultNotFound(ctx *Context) {
	if ctx.WantsJSON() {
		ctx.Response().Header().Set("Content-Type", "application/json")
	} else {
		ctx.Response().Header().Set("Content-Type", "text/html")
	}
	ctx.Status(http.StatusNotFound)

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
		err := ctx.View("errors/404")
		if err != nil {
			ctx.Error(err)
		}
	} else {
		ctx.WriteHTML("<!DOCTYPE html><html><head><title>404 Not Found</title></head><body>404 Not Found</body></html>")
	}
}
