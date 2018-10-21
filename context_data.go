package kira

// SetData ...
func (c *Context) SetData(key string, data interface{}) {
	c.data[key] = data
}

// GetData ...
func (c *Context) GetData(key string) interface{} {
	return c.data[key]
}
