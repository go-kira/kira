package session

import (
	"github.com/Lafriakh/log"

	"github.com/google/uuid"
)

// Store for manage the session
type Store struct {
	id      string // session id
	Handler Handler
	Values  map[interface{}]interface{}
}

// NewStore ...
func NewStore(name string, handler Handler) *Store {
	store := &Store{}
	store.SetID("") // session id
	store.Handler = handler
	store.Values = make(SKV)

	//  return the store instance
	return store
}

// Start the session, reading the data from a handler.
func (s *Store) Start() {
	s.loadSession()
}
func (s *Store) loadSession() {
	s.Values = s.ReadFromHandler()
}

// ReadFromHandler read the data from the handler
func (s *Store) ReadFromHandler() SKV {
	readedData, err := s.Handler.Read(s.GetID())
	if err != nil {
		log.Panic(err)
		return nil
	}

	// decode the session data
	var data SKV

	// if there no session values, return empty data with no errors
	if len(readedData) == 0 {
		return make(SKV)
	}

	// decode the session data
	err = DecodeGob(readedData, &data)
	if err != nil {
		log.Errorf("session decode: %v", err)
		return make(SKV)
	}

	return data
}

// SetID to set the id for the session or use the given id
func (s *Store) SetID(id string) {
	if id == "" {
		// generate the session id
		id = s.generateID()
	}
	// use the given session id
	s.id = id
}

// GetID - return the session id
func (s *Store) GetID() string {
	return s.id
}

// Save - write the data by the handler
func (s *Store) Save() {
	s.AgeFlashData()

	// encode the data from the handler
	encData, err := EncodeGob(s.Values)
	if err != nil {
		log.Panic(err)
	}

	s.Handler.Write(s.GetID(), encData)
}

// DestroyFromHandler - destroy the session
func (s *Store) DestroyFromHandler(id string) error {
	return s.Handler.Destroy(id)
}

// Put - put a key / value pair in the session.
func (s *Store) Put(key interface{}, value interface{}) {
	s.Values[key] = value
}

// Push - push a value onto a session slice.
func (s *Store) Push(key interface{}, value interface{}) {
	// init an empty slice
	var slice []interface{}
	// append any data to the slice if exists.
	slice = s.GetWithDefault(key, slice).([]interface{})
	// append the flash data.
	slice = append(slice, value)

	// puth the flash data to the session.
	s.Put(key, slice)
}

// Get - get a key / value pair from the session.
func (s *Store) Get(key interface{}) interface{} {
	// read data from the file
	return s.Values[key]
}

// GetWithDefault - get a key / value pair from the session.
func (s *Store) GetWithDefault(key interface{}, def interface{}) interface{} {
	// if there no value.
	if s.Values[key] == nil {
		return def
	}

	// read data from the file
	return s.Values[key]
}

// Remove - remove the key / value pair from the session.
func (s *Store) Remove(keys ...interface{}) {
	for _, val := range keys {
		// read data from the file
		delete(s.Values, val)
	}
}

// Flush - remove all of the items from the session.
func (s *Store) Flush() {
	// remove the key from session
	s.Values = make(SKV)
}

// Has - Checks if a key is present and not nil.
func (s *Store) Has(key interface{}) bool {
	_, ok := s.Values[key]
	if !ok {
		return false
	}

	return true
}

// Flash - a key / value pair to the session.
func (s *Store) Flash(key interface{}, value interface{}) {
	s.Put(key, value)

	s.Push("_flash.new", key)
}

// FlashPush - a key / value pair to the session.
func (s *Store) FlashPush(key interface{}, value interface{}) {
	s.Push(key, value)

	s.Push("_flash.new", key)
}

// AgeFlashData - Age the flash data for the session.
func (s *Store) AgeFlashData() {
	oldFlashes := s.GetWithDefault("_flash.old", []interface{}{}).([]interface{})
	newFlashes := s.GetWithDefault("_flash.new", []interface{}{}).([]interface{})
	// remove old flashes
	s.Remove(oldFlashes...)
	// put new flashes into old flashes.
	s.Put("_flash.old", newFlashes)
	// make new flashes empty
	s.Put("_flash.new", []interface{}{})
}

// Migrate - generate a new session ID for the session.
func (s *Store) Migrate(empty bool) {
	if empty {
		s.Values = make(SKV)
	}

	// remove old data from handler
	s.DestroyFromHandler(s.id)

	// generate new session id
	s.SetID(s.generateID())
}

func (s *Store) generateID() string {
	return uuid.New().String()
}

// GC - Garbage collector, to remove old session.
func (s *Store) GC() {
	s.Handler.GC()
}
