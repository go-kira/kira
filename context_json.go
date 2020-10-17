package kira

import (
	"encoding/json"
)

// JSON - Send response as json.
func (c *Context) JSON(data interface{}, code ...int) {
	c.Response().Header().Set("Content-Type", "application/json")

	// Status statusCode
	if len(code) > 0 {
		c.Status(code[0])
	}

	// Encode data
	if err := json.NewEncoder(c.Response()).Encode(data); err != nil {
		c.Error(err)
	}
}

// WantsJSON - validate if the request wants a json response.
func (c *Context) WantsJSON() bool {
	return c.Request().Header.Get("Accept") == "application/json"
}

// DecodeJSON - convert json from request body to interface.
func (c *Context) DecodeJSON(dst interface{}) error {
	return json.NewDecoder(c.Request().Body).Decode(dst)
}
