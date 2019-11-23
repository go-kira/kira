package kira

import (
	"fmt"
	"io"
	"net/http"
)

// Write writes the slice of bytes as an HTTP reply.
func (c *Context) Write(b []byte) {
	_, err := c.Response().Write(b)
	if err != nil {
		c.Error(err)
	}
}

// WriteString writes the content of s to the request response.
func (c *Context) WriteString(s string) {
	_, err := io.WriteString(c.Response(), s)
	if err != nil {
		c.Error(err)
	}
}

// WriteStringf formats according to a format specifier and writes the resulting to the request response.
func (c *Context) WriteStringf(format string, a ...interface{}) {
	fmt.Fprintf(c.Response(), format, a...)
}

// WriteHTML send content of html with the HTTP reply as an HTML content.
func (c *Context) WriteHTML(html string) {
	c.Response().Header().Set("Content-Type", "text/html")
	_, err := io.WriteString(c.Response(), html)
	if err != nil {
		c.Error(err)
	}
}

// WriteStatus Write HTTP header to the response and also write the status message to the body.
func (c *Context) WriteStatus(code int) {
	c.Response().WriteHeader(code)

	fmt.Fprint(c.Response(), http.StatusText(code))
}

// Status send a specific status with the HTTP reply.
func (c *Context) Status(code int) {
	c.Response().WriteHeader(code)
}

// File send a file with the HTTP reply.
func (c *Context) File(name string) {
	http.ServeFile(c.Response(), c.Request(), name)
}
