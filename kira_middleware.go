package kira

import (
	"net/http"
)

// Middleware interface
type Middleware interface {
	Handler(next http.Handler) http.Handler
}

// UseMiddlewares - assign many middlewares
func (a *App) UseMiddlewares(m []Middleware) {
	for _, middlware := range m {
		a.Middlewares = append(a.Middlewares, middlware)
	}
	// a.Middlewares = m
}

// UseMiddleware - add the middleware
func (a *App) UseMiddleware(m Middleware) {
	a.Middlewares = append(a.Middlewares, m)
}
