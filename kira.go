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
func (a *App) Run(addr ...string) *App {
	fmt.Printf("%v", hero)

	// Register the application routes
	a.RegisterRoutes()

	// Timezone
	tz := a.Configs.GetString("app.timezone")
	if tz != "" {
		// Now the framework will parse all the times in the given Timezone.
		os.Setenv("TZ", tz)
	}

	// TCP address
	serverAddr := serverAddr(a.Configs, addr...)
	if !a.Configs.GetBool("server.tls", false) {
		a.StartServer(serverAddr)
	} else {
		a.StartTLSServer(serverAddr)
	}

	// App instance
	return a
}

// NotFound custom not found handler.
func (a *App) NotFound(ctx HandlerFunc) {
	a.NotFoundHandler = ctx
}

// default not found handler.
func defaultNotFound(ctx *Context) {
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
		ctx.String("<!DOCTYPE html><html><head><title>404 Not Found</title></head><body>404 Not Found</body></html>")
	}
}
