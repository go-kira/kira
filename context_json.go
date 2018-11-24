package kira

import "encoding/json"

// JSON - Send response as json.
func (c *Context) JSON(data interface{}) {
	c.Response().Header().Set("Content-Type", "application/json")

	// Encode data
	if err := json.NewEncoder(c.Response()).Encode(data); err != nil {
		c.Error(err)
	}
}

// WantsJSON - validate if the request wants a json response.
func (c *Context) WantsJSON() bool {
	if c.Request().Header.Get("Accept") == "application/json" {
		return true
	}

	return false
}

// ParseJSON - convert json from request body to interface.
func (c *Context) ParseJSON(dst interface{}) {
	err := json.NewDecoder(c.Request().Body).Decode(dst)
	if err != nil {
		c.Error(err)
	}
}
