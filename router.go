package kira

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// Route struct
type Route struct {
	Name        string
	Methods     []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// SetName set the route name.
func (r *Route) SetName(name string) *Route {
	r.Name = name
	return r
}

// Middleware - set a middleware to the route.
func (r *Route) Middleware(middleware Middleware) *Route {
	r.HandlerFunc = middleware.Handler(r.HandlerFunc).ServeHTTP

	return r
}

// NewRouter return all routes.
func (a *App) NewRouter() *mux.Router {
	for _, route := range a.Routes {
		var handler http.Handler

		handler = route.HandlerFunc

		// append middlewares.
		for _, middleware := range a.Middlewares {
			handler = middleware.Handler(handler)
		}

		a.Router.Methods(route.Methods...).Path(route.Pattern).Name(route.Name).Handler(handler)
	}

	// 404 pages.
	a.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := NewContext(w, r, a.Log)

		if ctx.WantsJSON() {
			ctx.Header(http.StatusNotFound)
			ctx.JSON(struct {
				Error   int    `json:"error"`
				Message string `json:"message"`
			}{http.StatusNotFound, "404 not found"})
			return
		}

		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	})

	// return router
	return a.Router
}

// Static ...
func (a *App) Static(path, url string) {
	a.Router.PathPrefix(url).Handler(
		http.StripPrefix(url,
			http.FileServer(http.Dir(path)),
		),
	)
}

// UseRoutes - assign the routes
func (a *App) UseRoutes(m []Route) {
	for _, route := range m {
		a.Routes = append(a.Routes, &route)
		// a.Routes[route.Pattern] = &route
	}
}

// UseRoute for append route to the routes
func (a *App) UseRoute(m Route) {
	a.Routes = append(a.Routes, &m)
	// a.Routes[m.Pattern] = &m
}

// Methods ...
func (a *App) Methods(methods []string, pattern string, handler http.HandlerFunc) *Route {
	route := &Route{Methods: methods, Pattern: pattern, HandlerFunc: handler}
	a.Routes = append(a.Routes, route)

	return route
}

// GET request
func (a *App) GET(pattern string, handler http.HandlerFunc) *Route {
	route := &Route{Methods: []string{"GET"}, Pattern: pattern, HandlerFunc: handler}
	a.Routes = append(a.Routes, route)

	return route
}

// POST request
func (a *App) POST(pattern string, handler http.HandlerFunc) *Route {
	route := &Route{Methods: []string{"POST"}, Pattern: pattern, HandlerFunc: handler}
	a.Routes = append(a.Routes, route)

	return route
}

// PUT request
func (a *App) PUT(pattern string, handler http.HandlerFunc) *Route {
	route := &Route{Methods: []string{"PUT"}, Pattern: pattern, HandlerFunc: handler}
	a.Routes = append(a.Routes, route)

	return route
}

// DELETE request
func (a *App) DELETE(pattern string, handler http.HandlerFunc) *Route {
	route := &Route{Methods: []string{"DELETE"}, Pattern: pattern, HandlerFunc: handler}
	a.Routes = append(a.Routes, route)

	return route
}

// HEAD request
func (a *App) HEAD(pattern string, handler http.HandlerFunc) *Route {
	route := &Route{Methods: []string{"HEAD"}, Pattern: pattern, HandlerFunc: handler}
	a.Routes = append(a.Routes, route)

	return route
}

// OPTIONS request
func (a *App) OPTIONS(pattern string, handler http.HandlerFunc) *Route {
	route := &Route{Methods: []string{"OPTIONS"}, Pattern: pattern, HandlerFunc: handler}
	a.Routes = append(a.Routes, route)

	return route
}

// PATCH request
func (a *App) PATCH(pattern string, handler http.HandlerFunc) *Route {
	route := &Route{Methods: []string{"PATCH"}, Pattern: pattern, HandlerFunc: handler}
	a.Routes = append(a.Routes, route)

	return route
}

// // PathPrefix
// func (a *App) PathPrefix(tpl string, handler http.HandlerFunc) *Route {
// 	hand := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		a.Router.PathPrefix(tpl).Handler(http.StripPrefix(tpl, http.HandlerFunc(handler)))
// 	})

// 	route := &Route{HandlerFunc: hand}
// 	a.Routes = append(a.Routes, route)

// 	return route
// }

// Redirect to a url
func (a *App) Redirect(w http.ResponseWriter, req *http.Request, url string) {
	http.Redirect(w, req, url, 302)
}

// RedirectWithCode to a url with code
func (a *App) RedirectWithCode(w http.ResponseWriter, req *http.Request, url string, code int) {
	http.Redirect(w, req, url, code)
}

// RedirectWithError to a url
func (a *App) RedirectWithError(w http.ResponseWriter, req *http.Request, url string, err error) {
	a.Session.Flash("error", err.Error())

	http.Redirect(w, req, url, 302)
}

// Abort ...
func (a *App) Abort(w http.ResponseWriter, req *http.Request, code int) {
	errorPage := "errors/" + strconv.Itoa(code)
	// check if the error page template exists
	if _, err := os.Stat(a.Configs.GetString("VIEWS_PREFIX") + errorPage + a.Configs.GetString("VIEWS_FILE_SUFFIX")); os.IsNotExist(err) {
		errorPage = "errors/500"
	}

	w.WriteHeader(code)
	a.View.Render(w, req, errorPage)
}

// Query - return a request query value.
func (a *App) Query(request *http.Request, param string) string {
	return request.URL.Query().Get(param)
}

// Var - return a request var value.
func (a *App) Var(request *http.Request, rvar string) string {
	return mux.Vars(request)[rvar]
}
