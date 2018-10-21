package requestid

import (
	"context"
	"net/http"

	"github.com/go-kira/kon"
	"github.com/google/uuid"
)

// RequestID struct
type RequestID struct{ HeaderName string }

// NewRequestID - new instance of RequestID.
func NewRequestID(config *kon.Kon) *RequestID {
	return &RequestID{
		HeaderName: config.GetString("SERVER_REQUEST_ID"),
	}
}

// Handler - middleware handler
func (rq *RequestID) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// request id context
		ctx := context.WithValue(r.Context(), rq.HeaderName, rq.random())

		// set header
		w.Header().Set(rq.HeaderName, rq.random())

		// next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// random return random string for request id
func (rq *RequestID) random() string {
	return uuid.New().String()
}
