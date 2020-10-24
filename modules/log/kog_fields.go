package log

type Fields map[string]interface{}

// Get field value by name.
func (f Fields) Get(name string) interface{} {
	return f[name]
}
