package csrf

import (
	"net/http"

	"github.com/go-kira/kira"
	"github.com/gorilla/csrf"
)

// CSRF Middelware
type CSRF struct {
	*kira.App
}

// NewCSRF ...
func NewCSRF(app *kira.App) *CSRF {
	return &CSRF{app}
}

// Handler ...
func (c *CSRF) Handler(next http.Handler) http.Handler {
	CSRF := csrf.Protect(
		[]byte(c.Configs.GetString("app.key")),
		csrf.FieldName(c.Configs.GetString("csrf.field_name", "_token")),
		csrf.CookieName(c.Configs.GetString("csrf.cookie_name", "kira_csrf")),
	)

	return CSRF(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Go to the next request
		next.ServeHTTP(w, r)
	}))
}
