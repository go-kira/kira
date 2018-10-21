package session

import (
	"encoding/gob"
	"net/http"
	"time"

	"github.com/Lafriakh/kira/helpers"
	"github.com/go-kira/kon"
)

func init() {
	gob.Register(SKV{})
	gob.Register([]interface{}{})
	gob.Register(map[string]string{})
	gob.Register(map[string]error{})
}

// SKV - key value type to use it in other handlers
type SKV map[interface{}]interface{}

// Session ...
type Session struct {
	Store   *Store
	Options Options
}

// NewSession ...
func NewSession(config *kon.Kon, handler Handler, name string) *Session {
	opt := prepareOptions(config)
	store := NewStore(name, handler)
	// resturn the session
	return &Session{
		Store:   store,
		Options: opt,
	}
}

// SetOptions - to set custom options
func (s *Session) SetOptions(options Options) {
	s.Options = options
}

// StartGC starts GC job in a certain period.
func (s *Session) StartGC(config *kon.Kon) {
	// log.Debug("Session GC")
	s.Store.GC()
	time.AfterFunc(time.Duration(config.GetInt("SESSION_COOKIE_LIFETIME"))*time.Second, func() {
		s.StartGC(config)
	})
}

// All return all session values
func (s *Session) All() SKV {
	return s.Store.Values
}

// Get - get a key / value pair from the session.
func (s *Session) Get(key interface{}) interface{} {
	return s.Store.Get(key)
}

// GetWithDefault - get a key / value pair from the session.
func (s *Session) GetWithDefault(key interface{}, def interface{}) interface{} {
	return s.Store.GetWithDefault(key, def)
}

// Put - put a key / value pair in the session.
func (s *Session) Put(key interface{}, value interface{}) {
	s.Store.Put(key, value)
}

// Push - push a value onto a session slice.
func (s *Session) Push(key interface{}, value interface{}) {
	s.Store.Push(key, value)
}

// Remove - remove session key.
func (s *Session) Remove(keys ...interface{}) {
	// remove the key from session
	s.Store.Remove(keys...)

	// rewrite the session values without this key / value pair.
	// s.Save()
}

// Has - Checks if a key is present and not nil.
func (s *Session) Has(key interface{}) bool {
	return s.Store.Has(key)
}

// Flush - remove all of the items from the session.
func (s *Session) Flush() {
	// remove the key from session
	s.Store.Flush()

	// rewrite the session values with empty values.
	// s.Save()
}

// Flash - a key / value pair to the session.
func (s *Session) Flash(key interface{}, value interface{}) {
	s.Store.Flash(key, value)
}

// FlashPush - a key / value pair to the session.
func (s *Session) FlashPush(key interface{}, value interface{}) {
	s.Store.FlashPush(key, value)
}

// Regenerate - generate a new session identifier.
func (s *Session) Regenerate(response http.ResponseWriter) {
	// update session id
	s.Store.Migrate(false)
	// update cookies
	helpers.SetCookie(http.Cookie{
		Name:   s.Options.Name,
		Value:  helpers.EncodeBase64([]byte(s.Store.GetID())),
		MaxAge: s.Options.Lifetime,
	}, response)
}

// RegenerateWithEmpty - generate a new session identifier.
func (s *Session) RegenerateWithEmpty(response http.ResponseWriter) {
	// update session id
	s.Store.Migrate(true)
	// update cookies
	helpers.SetCookie(http.Cookie{
		Name:   s.Options.Name,
		Value:  helpers.EncodeBase64([]byte(s.Store.GetID())),
		MaxAge: s.Options.Lifetime,
	}, response)
}
