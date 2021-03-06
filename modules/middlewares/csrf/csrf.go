// We will re-write this based on this resource: https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/09.1.html
package csrf

import (
	"net/http"

	"github.com/go-kira/kira"
	"github.com/gorilla/csrf"
)

// New ...
func New() *CSRF {
	return &CSRF{}
}

// CSRF Middelware
type CSRF struct{}

// Middleware handler.
func (c *CSRF) Middleware(ctx *kira.Context, next kira.HandlerFunc) {
	// Here we convert the next context handler to the normal http.Handler.
	// We just wrap it so we can use it later with Gorilla CSRF middleware.
	var handler http.Handler
	// 3. Continue to the next handler
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx.SetRequest(r)
		ctx.SetResponse(w)

		// Continue
		next(ctx)
	})

	// 2. Set the token in the header.
	handler = func(n http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := csrf.Token(r)

			// Save the csrf token into the header.
			w.Header().Set(ctx.Config().GetString("csrf.header_name", "X-CSRF-Token"), token)

			// Save the csrf token into the context.
			ctx.SetData("csrf_token", token)

			// Move to the next request.
			n.ServeHTTP(w, r)
		})
	}(handler)

	// 1. Run the csrf middleware.
	handler = csrf.Protect(
		[]byte(ctx.Config().GetString("app.key")),
		csrf.FieldName(ctx.Config().GetString("csrf.field_name", "_token")),
		csrf.RequestHeader(ctx.Config().GetString("csrf.header_name", "X-CSRF-Token")),
		csrf.CookieName(ctx.Config().GetString("csrf.cookie_name", "kira_csrf")),
		csrf.Secure(ctx.Config().GetBool("csrf.secure", true)),
	)(handler)

	// Excute the middleware
	handler.ServeHTTP(ctx.Response(), ctx.Request())
}
