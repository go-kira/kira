package session

import (
	"sync"
	"time"
)

// MemoryHandler ...
type MemoryHandler struct {
	data     map[string][]byte
	Lifetime int
	lock     sync.RWMutex
}

// NewMemoryHandler return MemoryHandler instance
func NewMemoryHandler(path string, lifetime int) *MemoryHandler {
	return &MemoryHandler{
		data:     make(map[string][]byte),
		Lifetime: lifetime,
	}
}

// Read ...
func (m *MemoryHandler) Read(id string) ([]byte, error) {

	// read the data from the file
	m.lock.RLock()
	defer m.lock.RUnlock()

	// return the raw data
	return m.data[id], nil
}

// Write ...
func (m *MemoryHandler) Write(id string, data []byte) error {
	// lock
	m.lock.Lock()
	defer m.lock.Unlock()

	m.data[id] = data

	return nil
}

// Destroy ...
func (m *MemoryHandler) Destroy(id string) error {
	return nil
}

// GC ...
func (m *MemoryHandler) GC(maxlifetime time.Time) {
	//
}
