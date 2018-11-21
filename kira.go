package kira

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-kira/kira/session"
	"github.com/go-kira/kira/validation"
	"github.com/go-kira/kog"
	"github.com/go-kira/kon" // "github.com/Lafriakh/env"
	"github.com/gorilla/mux"
)

var hero = `
   __ __   _             
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

// you can customize this before init.
var (
	PathStore    = "./storage"
	PathApp      = PathStore + "/app"
	PathResource = "./resources"
	PathView     = PathResource + "/views"
	PathSass     = PathResource + "/sass"
	PathJS       = PathResource + "/javascript"
	PathSession  = PathStore + "/framework/sessions"
	PathLogs     = PathStore + "/framework/logs"
)

// App hold the framework options
type App struct {
	Routes      []*Route
	Middlewares []Middleware
	Router      *mux.Router
	View        View
	Validation  *validation.Validation
	Session     *session.Session
	Log         *kog.Logger
	Configs     *kon.Kon
	Env         string
	DB          *sql.DB

	isTLS bool
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

	// init view with app instance
	app.View.Data = make(map[string]interface{})
	app.View.App = app

	// validation
	app.Validation = validation.New()

	// session
	app.Session = setupSession(app.Configs)

	// define a Router
	app.Router = mux.NewRouter().StrictSlash(true)

	// return App instance
	return app
}

// Run the framework
func (a *App) Run() *App {
	fmt.Printf("%v", hero)

	// parse routes & middlewares
	a.NewRouter()

	// validate if the server need tls connection.
	if !a.isTLS {
		// Start the server
		a.StartServer()
	} else {
		// TLS
		a.StartTLSServer()
	}

	// return App instance
	return a
}

// getEnv for set the framework environment.
func getEnv() string {
	// Get the environment from .kira_env file.
	if _, err := os.Stat("./.kira_env"); !os.IsNotExist(err) {
		// path/to/whatever exists
		kiraEnv, err := ioutil.ReadFile("./.kira_env")
		if err != nil {
			kog.Panic(err)
		}
		return strings.TrimSpace(string(kiraEnv))
	}

	// Get the environment from system variable
	osEnv := os.Getenv("KIRA_ENV")
	if osEnv == "" {
		return "development"
	}
	return osEnv
}

func getConfig() *kon.Kon {
	var files = []string{"./config/application.yaml"}
	var env = fmt.Sprintf("./config/environments/%s.yaml", getEnv())

	if _, err := os.Stat(env); !os.IsNotExist(err) {
		files = append(files, env)
	}

	return kon.NewFromFile(files...)
}
