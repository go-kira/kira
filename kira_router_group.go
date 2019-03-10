package kira

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Group represent routes group.
type Group struct {
	app    *App
	prefix string
}

// Group adds a prefix to the given routes.
func (app *App) Group(prefix string, group func(Group)) {
	g := Group{
		app:    app,
		prefix: prefix,
	}

	// Register the group routes.
	group(g)
}

func (g Group) path(path string) string {
	return httprouter.CleanPath(g.prefix + httprouter.CleanPath(path))
}

// Get is a shortcut for app.Get with the group prefix.
func (g Group) Get(path string, handler HandlerFunc, middlewares ...Middleware) {
	g.app.Get(g.path(path), handler, middlewares...)
}

// Head is a shortcut for app.Head with the group prefix.
func (g Group) Head(path string, handler HandlerFunc, middlewares ...Middleware) {
	g.app.Head(g.path(path), handler, middlewares...)
}

// Post is a shortcut for app.Post with the group prefix.
func (g Group) Post(path string, handler HandlerFunc, middlewares ...Middleware) {
	g.app.Post(g.path(path), handler, middlewares...)
}

// Put is a shortcut for app.Put with the group prefix.
func (g Group) Put(path string, handler HandlerFunc, middlewares ...Middleware) {
	g.app.Put(g.path(path), handler, middlewares...)
}

// Patch is a shortcut for app.Patch with the group prefix.
func (g Group) Patch(path string, handler HandlerFunc, middlewares ...Middleware) {
	g.app.Patch(g.path(path), handler, middlewares...)
}

// Delete is a shortcut for app.Delete with the group prefix.
func (g Group) Delete(path string, handler HandlerFunc, middlewares ...Middleware) {
	g.app.Delete(g.path(path), handler, middlewares...)
}

// Options is a shortcut for app.Delete with the group prefix.
func (g Group) Options(path string, handler HandlerFunc, middlewares ...Middleware) {
	g.app.Options(g.path(path), handler, middlewares...)
}

// ServeFiles is a shortcut for app.ServeFiles with the group prefix.
func (g Group) ServeFiles(path string, root http.FileSystem) {
	g.app.ServeFiles(g.path(path), root)
}
