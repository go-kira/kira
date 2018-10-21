package kira

import (
	"net/http"
)

// Validate - for request validation.
func (app *App) Validate(request *http.Request, fields map[string]string) map[string]string {
	var errs []error
	errorMap := make(map[string]string)
	for key, rules := range fields {
		value := request.FormValue(key)

		if err := app.Validation.Validate(key, value, rules); err != nil {
			// append the error to the slice errs
			errs = append(errs, err)
			errorMap[key] = err.Error()
		}
	}

	if len(errs) > 0 {
		for key, err := range errorMap {
			app.Session.FlashPush("errors."+key, err)
		}
		app.Session.Flash("errors", errorMap)

		return errorMap
	}
	return nil
}
