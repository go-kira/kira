package kira

import "fmt"

// Panic - stop the request with panic
func (c *Context) Panic(msg ...interface{}) {
	// Just panic and the recover will come to save us :)
	// c.Log().Panic(fmt.Sprint(msg...))
	panic(fmt.Sprint(msg...))
}

// Error - stop the request with error message
func (c *Context) Error(msg ...interface{}) {
	// Just panic and the recover will come to save us :)
	c.Log().Error(fmt.Sprint(msg...))
}
