package limitbody

import (
	"net/http"

	"github.com/go-kira/kira"
)

// MB - one MB.
const MB = 1 << 20

// Limitbody - Middleware.
type Limitbody struct {
	*kira.App
}

// Newlimitbody - return Limitbody instance
func Newlimitbody(app *kira.App) *Limitbody {
	return &Limitbody{app}
}

// Handler - middelware handler
func (l *Limitbody) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength > l.App.Configs.GetInt64("SERVER_BODY_LIMIT")*MB {
			http.Error(w, "Request too large", http.StatusExpectationFailed)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, l.App.Configs.GetInt64("SERVER_BODY_LIMIT")*MB)

		next.ServeHTTP(w, r)
	})
}
