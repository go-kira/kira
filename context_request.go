package kira

import (
	"github.com/julienschmidt/httprouter"
)

// Query get request query value
func (c *Context) Query(param string) string {
	return c.request.URL.Query().Get(param)
}

// Param is an alias of var method.
func (c *Context) Param(param string) string {
	params := httprouter.ParamsFromContext(c.Request().Context())

	return params.ByName(param)
}
