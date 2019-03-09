package kira

// Middleware interface
type Middleware interface {
	// Name() string
	Middleware(*Context, HandlerFunc)
}

// Middleware - add the middleware
func (app *App) Middleware(middlewares ...Middleware) {
	for _, m := range middlewares {
		app.Middlewares = append(app.Middlewares, m)
	}
}

// Use is an alias of Middleware method.
func (app *App) Use(middlewares ...Middleware) {
	app.Middleware(middlewares...)
}
