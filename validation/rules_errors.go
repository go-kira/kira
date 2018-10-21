package validation

const (
	errRequired = "The %s field is required."
	errInteger  = "The %s field must contain an integer."
	errNumeric  = "The %s field must contain only numbers."
	errMax      = "The %s field cannot exceed %v characters in length."
	errMin      = "The %s field must be at least %v characters in length."
	errBetween  = "the given value is %v, is not between %v, %v."
	errEmail    = "The %s field must contain a valid email address."
	errURL      = "The %s field must contain a valid url."
)
