package kira

import (
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

// String writes the content of s to the request response.
func (c *Context) String(s string) {
	_, err := io.WriteString(c.Response(), s)
	if err != nil {
		c.Error(err)
	}
}

// File send a file with the HTTP reply.
func (c *Context) File(name string) {
	http.ServeFile(c.Response(), c.Request(), name)
}
