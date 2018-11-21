package kira

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Query ...
func (c *Context) Query(param string) string {
	return c.request.URL.Query().Get(param)
}

// Var - get route variable.
func (c *Context) Var(variable string) string {
	vars := mux.Vars(c.request)

	if val, ok := vars[variable]; ok {
		return val
	}

	return ""
}

// Handle Handler any request type.
func (a *App) Handle(pattern string, ctx ContextFunc) *Route {
	route := &Route{Methods: []string{}, Pattern: pattern, HandlerFunc: func(w http.ResponseWriter, req *http.Request) {
		ctx(NewContext(w, req, a.Log))
	}}
	a.Routes = append(a.Routes, route)

	return route
}

// Get request
func (a *App) Get(pattern string, ctx ContextFunc) *Route {
	route := &Route{Methods: []string{"GET"}, Pattern: pattern, HandlerFunc: func(w http.ResponseWriter, req *http.Request) {
		ctx(NewContext(w, req, a.Log))
	}}
	a.Routes = append(a.Routes, route)

	return route
}

// Post request
func (a *App) Post(pattern string, ctx ContextFunc) *Route {
	route := &Route{Methods: []string{"POST"}, Pattern: pattern, HandlerFunc: func(w http.ResponseWriter, req *http.Request) {
		ctx(NewContext(w, req, a.Log))
	}}
	a.Routes = append(a.Routes, route)

	return route
}

// Put request
func (a *App) Put(pattern string, ctx ContextFunc) *Route {
	route := &Route{Methods: []string{"PUT"}, Pattern: pattern, HandlerFunc: func(w http.ResponseWriter, req *http.Request) {
		ctx(NewContext(w, req, a.Log))
	}}
	a.Routes = append(a.Routes, route)

	return route
}

// Delete request
func (a *App) Delete(pattern string, ctx ContextFunc) *Route {
	route := &Route{Methods: []string{"DELETE"}, Pattern: pattern, HandlerFunc: func(w http.ResponseWriter, req *http.Request) {
		ctx(NewContext(w, req, a.Log))
	}}
	a.Routes = append(a.Routes, route)

	return route
}

// Head request
func (a *App) Head(pattern string, ctx ContextFunc) *Route {
	route := &Route{Methods: []string{"HEAD"}, Pattern: pattern, HandlerFunc: func(w http.ResponseWriter, req *http.Request) {
		ctx(NewContext(w, req, a.Log))
	}}
	a.Routes = append(a.Routes, route)

	return route
}

// Options request
func (a *App) Options(pattern string, ctx ContextFunc) *Route {
	route := &Route{Methods: []string{"OPTIONS"}, Pattern: pattern, HandlerFunc: func(w http.ResponseWriter, req *http.Request) {
		ctx(NewContext(w, req, a.Log))
	}}
	a.Routes = append(a.Routes, route)

	return route
}

// Patch request
func (a *App) Patch(pattern string, ctx ContextFunc) *Route {
	route := &Route{Methods: []string{"PATCH"}, Pattern: pattern, HandlerFunc: func(w http.ResponseWriter, req *http.Request) {
		ctx(NewContext(w, req, a.Log))
	}}
	a.Routes = append(a.Routes, route)

	return route
}
