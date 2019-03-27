package kira

// SetData ...
func (c *Context) SetData(key string, data interface{}) {
	c.data[key] = data
}

// GetData ...
func (c *Context) GetData(key string) interface{} {
	return c.data[key]
}

// HasData ...
func (c *Context) HasData(key string) bool {
	if _, ok := c.data[key]; ok {
		return true
	}
	return false
}
