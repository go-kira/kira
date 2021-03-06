package kira

import (
	"github.com/julienschmidt/httprouter"
)

// Query get request query value
func (c *Context) Query(param string) string {
	return c.request.URL.Query().Get(param)
}

// HasQuery checks if the request has the given query.
func (c *Context) HasQuery(key string) bool {
	v := c.request.URL.Query()
	if v == nil {
		return false
	}

	// Check if the query exists in the url.
	if _, ok := v[key]; ok {
		return true
	}
	return false
}

// Param is an alias of var method.
func (c *Context) Param(param string) string {
	return httprouter.ParamsFromContext(c.Request().Context()).ByName(param)
}

// ParseMultipartForm pars
func (c *Context) ParseMultipartForm() error {
	return c.Request().ParseMultipartForm(c.Config().GetInt64("server.body_limit", 32) * MB)
}
