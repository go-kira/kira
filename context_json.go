package kira

// WantsJSON - validate if the request wants a json response.
func (c *Context) WantsJSON() bool {
	if c.Request().Header.Get("Accept") == "application/json" {
		return true
	}

	return false
}
