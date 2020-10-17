package kira

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Route represent a route.
type Route struct {
	Method      string
	Path        string
	HandlerFunc HandlerFunc
	Middlewares []Middleware
}

// Middleware add a middleware to the route.
func (r *Route) Middleware(midd ...Middleware) {
	r.Middlewares = append(r.Middlewares, midd...)
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
		app.Router.Handler(
			// Method
			route.Method,
			// Path
			route.Path,
			// Handler
			buildRoute(app, route.HandlerFunc, route.Middlewares),
		)
	}

	// 404
	if app.NotFoundHandler == nil {
		app.Router.NotFound = buildRoute(app, defaultNotFound, nil)
	} else {
		app.Router.NotFound = buildRoute(app, app.NotFoundHandler, nil)
	}

	return app.Router
}

// Change the middleware to support middleware chain.
// This function will take the middleware and the next handler as a parameters.
// Then return a handler that accept the next handler as a parameter.
func buildMiddleware(middleware Middleware, next HandlerFunc) HandlerFunc {
	return func(ctx *Context) {
		middleware.Middleware(ctx, next)
	}
}

// buildRoute create the context for the route and attach the middlewares to it if exists.
func buildRoute(app *App, handler HandlerFunc, rm []Middleware) http.HandlerFunc {
	// Route middlewares
	if len(rm) > 0 {
		for _, m := range rm {
			handler = buildMiddleware(m, handler)
		}
	}

	// Assign default middlewares to all handlers.
	for _, defaultMiddleware := range defaultMiddlewares() {
		handler = buildMiddleware(defaultMiddleware, handler)
	}

	// Global Middlewares
	if len(app.Middlewares) > 0 {
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
			handler = buildMiddleware(m, handler)
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Root context.
		// TODO
		//  - Set default values in the context.
		//  - Like request id, csrf...
		ctx := app.pool.Get().(*Context)
		ctx.response = w
		ctx.request = r

		// Run the chain
		handler(ctx)

		// Release the pool
		contextPool.Put(ctx)
	}
}

// create new route instance.
func createRoute(app *App, method string, path string, handler HandlerFunc, middlewares ...Middleware) *Route {
	route := &Route{
		Method:      method,
		Path:        path,
		Middlewares: middlewares,
		HandlerFunc: handler,
	}

	// Append the route
	app.Routes = append(app.Routes, route)

	return route
}

// Get Handle GET requests.
func (app *App) Get(path string, ctx HandlerFunc, middlewares ...Middleware) *Route {
	return createRoute(app, "GET", path, ctx, middlewares...)
}

// Head Handle HEAD requests.
func (app *App) Head(path string, ctx HandlerFunc, middlewares ...Middleware) *Route {
	return createRoute(app, "HEAD", path, ctx, middlewares...)
}

// Post Handle POST requests.
func (app *App) Post(path string, ctx HandlerFunc, middlewares ...Middleware) *Route {
	return createRoute(app, "POST", path, ctx, middlewares...)
}

// Put Handle PUT requests.
func (app *App) Put(path string, ctx HandlerFunc, middlewares ...Middleware) *Route {
	return createRoute(app, "PUT", path, ctx, middlewares...)
}

// Patch Handle PATCH requests.
func (app *App) Patch(path string, ctx HandlerFunc, middlewares ...Middleware) *Route {
	return createRoute(app, "PATCH", path, ctx, middlewares...)
}

// Delete Handle DELETE requests.
func (app *App) Delete(path string, ctx HandlerFunc, middlewares ...Middleware) *Route {
	return createRoute(app, "DELETE", path, ctx, middlewares...)
}

// Options Handle OPTIONS requests.
func (app *App) Options(path string, ctx HandlerFunc, middlewares ...Middleware) *Route {
	return createRoute(app, "OPTIONS", path, ctx, middlewares...)
}

// ServeFiles serve files in the given root.
func (app *App) ServeFiles(path string, root http.FileSystem) {
	app.Router.ServeFiles(path, root)
}
