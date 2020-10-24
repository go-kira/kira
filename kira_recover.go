package kira

import (
	"fmt"
	"net/http"
	"runtime"
)

// ErrorFrame ...
type ErrorFrame struct {
	File string `json:"file"`
	Func string `json:"func"`
	Line int    `json:"line"`
}

// ErrorJSON ...
type ErrorJSON struct {
	Message string       `json:"message"`
	Frames  []ErrorFrame `json:"frames,omitempty"`
}

// Middleware handler.
func defaultPanic(ctx *Context, err interface{}) {
	// log the error
	ctx.Log().Errorf("%s", err)

	// Headers
	if ctx.WantsJSON() {
		ctx.Response().Header().Set("Content-Type", "application/json")
	} else {
		ctx.Response().Header().Set("Content-Type", "text/html")
	}
	ctx.Status(http.StatusInternalServerError)

	// if the debug mode is enabled, add the stack to the error view
	if ctx.Config().GetBool("app.debug", false) {
		if ctx.WantsJSON() { // JSON
			var frames []ErrorFrame
			for _, frame := range getFrames(100) {
				frames = append(frames, ErrorFrame{
					File: frame.File,
					Func: frame.Func.Name(),
					Line: frame.Line,
				})
			}
			ctx.JSON(ErrorJSON{
				Message: fmt.Sprintf("%s", err),
				Frames:  frames,
			})
		} else { // HTML
			if ctx.ViewExists("error/debug") {
				ctx.View("errors/debug", Map{
					"message": err,
					"frames":  getFrames(100),
				})
			} else {
				ctx.WriteString("<p>We're sorry, but something went wrong.</p> \n\n")
				ctx.WriteStringf("<p>Message: <strong>%s</strong></p>\nFrames:\n\n", err)
				for _, frame := range getFrames(100) {
					ctx.WriteStringf("<pre>Func: %s \nFile: %s \nLine: %d</pre>\n\n", frame.Func.Name(), frame.File, frame.Line)
				}
			}
		}

		return
	}

	// Normal mode
	if ctx.WantsJSON() {
		ctx.JSON(ErrorJSON{
			Message: fmt.Sprintf("%s", err),
		})
	} else {
		if ctx.ViewExists("errors/500") {
			ctx.View("errors/500")
		} else {
			ctx.WriteString(`<html><head><title>Internal Server Error</title></head><body>We're sorry, but something went wrong.</body></html>`)
		}
	}
}

func getFrames(limit int) (framesSlice []runtime.Frame) {
	// Ask runtime.Callers for up to 10 pcs, including runtime.Callers itself.
	pc := make([]uintptr, limit)
	// TODO: later we need to hide unnecessary callers.
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
