package logger

import (
	"net/http"
	"time"

	"github.com/Lafriakh/kira"
	"github.com/go-kira/kog"
)

const basePath = "storage/logs/"

// colors
// var (
// yellow = color.New(color.FgYellow).SprintFunc()
// red   = color.New(color.FgRed).SprintFunc()
// green = color.New(color.FgGreen).SprintFunc()
// blue  = color.New(color.FgBlue).SprintFunc()
// cyan  = color.New(color.FgCyan).SprintFunc()
// )

// Log - log middleware
type Log struct {
	*kira.App
}

// NewLogger ...
func NewLogger(app *kira.App) *Log {
	return &Log{app}
}

// Handler - middleware handler
func (l *Log) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start time
		var now = time.Now()
		// Store the status code
		statusRecorder := &statusRecorder{w, http.StatusOK}

		// Run the request
		next.ServeHTTP(statusRecorder, r)

		// logger message
		// [INFO] [5b15c100-fdd3-482f-ad0c-037d4159d066] 2018/10/17 00:17:26 | [500] GET /seasons?page=1 63.803024ms
		l.App.Log.Infof("[%s] %s [%d] %s %s %v",
			// request id
			r.Context().Value(l.App.Configs.GetString("SERVER_REQUEST_ID")).(string),
			// time
			kog.FormatTime(time.Now()),
			// status code
			statusRecorder.statusCode,
			// method
			r.Method,
			// request path
			r.RequestURI,
			// request duration
			time.Since(now),
		)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader - store the header to use it later.
func (r *statusRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}
