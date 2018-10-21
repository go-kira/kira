package example

import (
	"github.com/Lafriakh/kira"
)

// Example - kira middleware example.
type Example struct {
	*kira.App
}

// NewExample - a new instance of Example
func NewExample(app *kira.App) *Example {
	return &Example{app}
}

// func (m *Example) ServeHTTP(next *http.Handler) *http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Before request
// 		next.ServeHTTP(w, r)
// 		// After request
// 	})
// }

// func (m *Example) ServeHTTP(next *kira.Context) *http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Before request
// 		next.ServeHTTP(w, r)
// 		// After request
// 	})
// }
