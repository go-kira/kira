package recover

import (
	"net/http"
	"runtime"
	"time"

	"github.com/Lafriakh/kira"
	"github.com/go-kira/kog"
)

// Recover - Middleware
type Recover struct {
	*kira.App
}

// NewRecover - return recover instance
func NewRecover(app *kira.App) *Recover {
	return &Recover{app}
}

// Handler - middelware handler
func (rc *Recover) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		defer func() {
			r := recover()
			// We have a problem here
			if r != nil {
				requestID := request.Context().Value(rc.App.Configs.GetString("SERVER_REQUEST_ID")).(string)

				// log the error
				rc.App.Log.Errorf("[%s] %s | %s", requestID, kog.FormatTime(time.Now()), r)

				// write header
				w.WriteHeader(http.StatusInternalServerError)

				// if the debug mode is enabled, add the stack to the error view
				if rc.App.Configs.GetBool("DEBUG") {
					rc.View.Data["message"] = r
					rc.View.Data["frames"] = getFrames(100)
					rc.View.Render(w, request, "errors/debug")
					return
				}

				// display error page
				rc.View.Render(w, request, "errors/500")
				return
			}
		}()

		next.ServeHTTP(w, request)
		return
	})
}

func getFrames(limit int) (framesSlice []runtime.Frame) {
	// Ask runtime.Callers for up to 10 pcs, including runtime.Callers itself.
	pc := make([]uintptr, limit)
	n := runtime.Callers(0, pc)
	if n == 0 {
		// No pcs available. Stop now.
		// This can happen if the first argument to runtime.Callers is large.
		return
	}

	pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
	frames := runtime.CallersFrames(pc)

	// Loop to get frames.
	// A fixed number of pcs can expand to an indefinite number of Frames.
	for {
		frame, more := frames.Next()
		framesSlice = append(framesSlice, frame)
		if !more {
			break
		}
	}

	return framesSlice
}
