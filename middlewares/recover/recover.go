package recover

import (
	"net/http"
	"runtime"

	"github.com/go-kira/kira"
)

// Recover - Middleware
type Recover struct{}

// New - return recover instance
func New() *Recover {
	return &Recover{}
}

// Middleware handler.
func (rc *Recover) Middleware(ctx *kira.Context, next kira.HandlerFunc) {
	defer func() {
		r := recover()
		// We have a problem here
		if r != nil {
			headerName := ctx.Config().GetString("server.request_id", "X-Request-Id")

			requestID := ctx.Request().Context().Value(headerName)

			// log the error
			ctx.Log().Errorf("%s %s", r, requestID)

			// write header
			ctx.HeaderStatus(http.StatusInternalServerError)

			// if the debug mode is enabled, add the stack to the error view
			if ctx.Config().GetBool("app.debug", false) {
				if ctx.ViewExists("error/debug") {
					ctx.View("errors/debug", kira.ViewData{
						"message": r,
						"frames":  getFrames(100),
					})
				} else {
					ctx.String("We're sorry, but something went wrong. \n\n")
					ctx.Stringf("Message: %s\nFrames:\n\n", r)
					for _, frame := range getFrames(100) {
						ctx.Stringf("Func: %s \nFile: %s \nLine: %d\n\n", frame.Func.Name(), frame.File, frame.Line)
					}
				}
				return
			}

			// display error page
			ctx.View("errors/500")
			return
		}
	}()

	next(ctx)
	return
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
