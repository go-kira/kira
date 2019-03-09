package kira

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Route represent a route.
type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
	Middlewares []Middleware
}

// Middleware add a middleware to the route.
func (r *Route) Middleware(midd ...Middleware) {
	for _, middleware := range midd {
		r.Middlewares = append(r.Middlewares, middleware)
	}
}

// Use is an alias of Middleware method.
func (r *Route) Use(midd ...Middleware) {
	r.Middleware(midd...)
}

// RegisterRoutes it's simply register the routes into the router.
func (app *App) RegisterRoutes() *httprouter.Router {
	// build the routes and attach the middlewares to every route.
	for _, route := range app.Routes {
		// Register the route.
		app.Router.Handler(route.Method, route.Path, route.HandlerFunc)
	}

	// 404
	if app.NotFoundHandler == nil {
		app.Router.NotFound = buildRoute(app, defaultNotFound, nil)
	} else {
		app.Router.NotFound = buildRoute(app, app.NotFoundHandler, nil)
	}

	return app.Router
}

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
	} else { // HTML
		// Validate if the template exists
		if ctx.ViewExists("errors/404") {
			ctx.View("errors/404")
		} else {
			ctx.String("<!DOCTYPE html><html><head><title>404 Not Found</title></head><body>404 Not Found</body></html>")
		}
	}
}

// buildRoute create the context for the route and attach the middlwares to it if exists.
func buildRoute(app *App, handler HandlerFunc, route *Route) http.HandlerFunc {
	// Change the middleware to support middleware chain.
	// This function will take the middleware and the next handler as a parameters.
	// Then return a handler that accept the next handler as a parameter.
	var middlewareHandler func(Middleware, HandlerFunc) HandlerFunc
	middlewareHandler = func(middleware Middleware, next HandlerFunc) HandlerFunc {
		return func(ctx *Context) {
			middleware.Middleware(ctx, next)
		}
	}

	// Route middlewares
	if route != nil && len(route.Middlewares) > 0 {
		for _, m := range route.Middlewares {
			handler = middlewareHandler(m, handler)
		}
	}

	// Global Middlewares
	for _, m := range app.Middlewares {
		// except := app.Configs.GetSliceString("excluded_middleware." + m.Name())
		//
		// // Move to the next router if the route is nil.
		// if route == nil || helpers.Contains(except, "*") {
		// 	continue
		// }
		//
		// if !helpers.Contains(except, route.Path) {
		// 	handler = middlewareHandler(m, handler)
		// }

		handler = middlewareHandler(m, handler)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Root context.
		// TODO
		//  - Set default values in the context.
		//  - Like request id, csrf...
		c := NewContext(w, r, app)

		// Run the chain
		handler(c)
	}
}

// create new route instance.
func createRoute(app *App, method string, path string, ctx HandlerFunc, middlewares ...Middleware) *Route {
	route := &Route{
		Method:      method,
		Path:        path,
		Middlewares: middlewares,
	}

	route.HandlerFunc = buildRoute(app, ctx, route)

	// Append the route
	app.Routes = append(app.Routes, route)

	return route
}

// Handle GET requests.
func (app *App) Get(path string, ctx HandlerFunc, middlewares ...Middleware) *Route {
	return createRoute(app, "GET", path, ctx, middlewares...)
}

// Handle HEAD requests.
func (app *App) Head(path string, ctx HandlerFunc, middlewares ...Middleware) *Route {
	return createRoute(app, "HEAD", path, ctx, middlewares...)
}

// Handle POST requests.
func (app *App) Post(path string, ctx HandlerFunc, middlewares ...Middleware) *Route {
	return createRoute(app, "POST", path, ctx, middlewares...)
}

// Handle PUT requests.
func (app *App) Put(path string, ctx HandlerFunc, middlewares ...Middleware) *Route {
	return createRoute(app, "PUT", path, ctx, middlewares...)
}

// Handle PATCH requests.
func (app *App) Patch(path string, ctx HandlerFunc, middlewares ...Middleware) *Route {
	return createRoute(app, "PATCH", path, ctx, middlewares...)
}

// Handle DELETE requests.
func (app *App) Delete(path string, ctx HandlerFunc, middlewares ...Middleware) *Route {
	return createRoute(app, "DELETE", path, ctx, middlewares...)
}

// Handle OPTIONS requests.
func (app *App) Options(path string, ctx HandlerFunc, middlewares ...Middleware) *Route {
	return createRoute(app, "OPTIONS", path, ctx, middlewares...)
}

// Handle ServeFiles requests.
func (app *App) ServeFiles(path string, root http.FileSystem) {
	app.Router.ServeFiles(path, root)
}
