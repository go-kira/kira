package session

// Handler ...
// the data will be in map[interface{}]interface{}
type Handler interface {
	Read(id string) ([]byte, error)
	Write(id string, data []byte) error
	Destroy(id string) error
	GC()
}
