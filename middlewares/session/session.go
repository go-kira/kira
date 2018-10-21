package session

import (
	"net/http"

	"github.com/Lafriakh/kira"
	"github.com/Lafriakh/kira/helpers"
	"github.com/Lafriakh/kira/session"
)

// Middleware - Middleware
type Middleware struct {
	*kira.App
}

// NewMiddleware - return session middlware instance
func NewMiddleware(app *kira.App) *Middleware {
	// start session GC
	app.Session.StartGC(app.Configs)

	// return session middleware
	return &Middleware{app}
}

// Handler - middelware handler
func (m *Middleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		store := m.startSession(request)

		// save cookies
		http.SetCookie(response, session.NewCookie(
			m.App.Configs,
			m.Session.Options.Name,
			helpers.EncodeBase64([]byte(store.GetID())),
			m.Session.Options,
		))

		next.ServeHTTP(response, request)

		// save the session
		store.Save()
	})
}

func (m *Middleware) startSession(req *http.Request) *session.Store {
	session := m.getSession(req)
	session.Start()
	return session
}

// getSession get  the session from the store
func (m *Middleware) getSession(request *http.Request) *session.Store {
	store := m.Session.Store

	// read the session id from the cookie
	cookie, _ := helpers.GetCookie(m.Session.Options.Name, request)
	store.SetID(session.ParseID(cookie))

	return store
}
