package csrf

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/Lafriakh/kira"
	"github.com/Lafriakh/kira/helpers"
	"github.com/Lafriakh/kira/session"
	"github.com/go-kira/kon"
	"golang.org/x/net/xsrftoken"
)

var (
	safeMethods = []string{"GET", "HEAD", "OPTIONS", "TRACE"}
)

// errors
var (
	// ErrNoToken is returned if no CSRF token is supplied in the request.
	ErrNoToken = errors.New("CSRF token not found in request")
	// ErrBadToken is returned if the CSRF token in the request does not match
	// the token in the session, or is otherwise malformed.
	ErrBadToken = errors.New("CSRF token invalid")
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if there a token, if not generate new one.
		if !c.App.Session.Has("_token") {
			c.RegenerateToken(c.App.Configs)
		}

		// context
		ctx := context.WithValue(r.Context(), "csrf", c.Token())

		// only on POST...
		if err := tokensMatch(c.App.Configs, r, c.App.Session.Get("_token").(string)); err == nil || isReading(r) {
			// addCookieToResponse
			c.addCookieToResponse(c.App.Configs, w, r)
		} else {
			log.Panic(err)
		}

		// Go to the next request
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Determine if the HTTP request uses a ‘read’ verb.
func isReading(r *http.Request) bool {
	if helpers.Contains(safeMethods, r.Method) {
		return true
	}
	return false
}

// Determine if the session and input CSRF tokens match.
func tokensMatch(config *kon.Kon, r *http.Request, sessionToken string) error {
	// get the token from the request.
	getToken := getTokenFromRequest(config, r)
	if getToken == "" {
		return ErrNoToken
	}

	// decode the token from session.
	token, err := helpers.DecodeURLBase64(getToken)
	if err != nil {
		return err
	}

	// validate if the token from the request equal to the token from the session.
	if string(token) == sessionToken {
		// validate the session timing
		if xsrftoken.Valid(string(token), config.GetString("APP_KEY"), "", "") {
			return nil
		}
	}

	return ErrBadToken
}

func getTokenFromRequest(config *kon.Kon, r *http.Request) string {
	// 1. Check the HTTP header first.
	token := r.Header.Get(config.GetString("CSRF_HEADER_NAME"))

	// 2. Fall back to the POST (form) value.
	if token == "" {
		token = r.PostFormValue(config.GetString("CSRF_FIELD_NAME"))
	}

	// 3. Finally, fall back to the multipart form (if set).
	if token == "" && r.MultipartForm != nil {
		vals := r.MultipartForm.Value[config.GetString("CSRF_FIELD_NAME")]

		if len(vals) > 0 {
			token = vals[0]
		}
	}

	return token
}

func (c *CSRF) addCookieToResponse(config *kon.Kon, w http.ResponseWriter, r *http.Request) {
	// set cookie to the response
	http.SetCookie(w, session.NewCookie(
		config,
		config.GetString("CSRF_COOKIE_NAME"),
		helpers.EncodeBase64([]byte(c.App.Session.Get("_token").(string))),
		c.App.Session.Options,
	))

}
