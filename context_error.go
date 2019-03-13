package kira

import "bytes"

var errSeparator = ":\n\t"

// Error represent kira error type.
type Error struct {
	Path string
	Op   string
	Err  error
}

func (e *Error) Error() string {
	b := new(bytes.Buffer)

	if e.Op != "" {
		pad(b, ": ")
		b.WriteString(string(e.Op))
	}
	if e.Path != "" {
		pad(b, ": ")
		b.WriteString(string(e.Path))
	}
	if e.Err != nil {
		if prevErr, ok := e.Err.(*Error); ok {
			if !prevErr.isZero() {
				pad(b, errSeparator)
				b.WriteString(e.Err.Error())
			}
		} else {
			pad(b, ": ")
			b.WriteString(e.Err.Error())
		}
	}
	if b.Len() == 0 {
		return "no error"
	}
	return b.String()
}

func (e *Error) isZero() bool {
	return e.Path == "" && e.Op == "" && e.Err == nil
}

// pad appends str to the buffer if the buffer already has some data.
func pad(b *bytes.Buffer, str string) {
	if b.Len() == 0 {
		return
	}
	b.WriteString(str)
}
