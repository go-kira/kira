package session

import (
	"bytes"
	"encoding/gob"
	"net/http"
	"time"

	"github.com/go-kira/kira/helpers"
	"github.com/go-kira/kon"
)

// EncodeGob the session data
func EncodeGob(input interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	encCache := gob.NewEncoder(buf)
	err := encCache.Encode(input)

	return buf.Bytes(), err
}

// DecodeGob the session data
func DecodeGob(input []byte, data interface{}) error {
	buf := bytes.NewBuffer(input)
	decCache := gob.NewDecoder(buf)
	err := decCache.Decode(data)

	return err
}

// GetMaxLifetime - this return the time type when the session expired.
func GetMaxLifetime(seconds int) time.Time {
	lifetime := time.Duration(seconds) * time.Second

	return time.Now().Add(lifetime)
}

// ParseID - take an id and convert it to base64.
func ParseID(id string) string {
	base64, err := helpers.DecodeBase64(id)
	if err != nil {
		return ""
	}
	return string(base64)
}

// NewCookie returns an http.Cookie with the options set. It also sets
// the Expires field calculated based on the MaxAge value, for Internet
// Explorer compatibility.
func NewCookie(config *kon.Kon, name, value string, options Options) *http.Cookie {
	// save cookie
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     options.Path,
		Domain:   options.Domain,
		MaxAge:   options.Lifetime,
		Secure:   options.Secure,
		HttpOnly: options.HTTPOnly,
	}

	if options.Lifetime > 0 {
		cookie.Expires = GetMaxLifetime(options.Lifetime)
	} else if options.Lifetime < 0 {
		// Set it to the past to expire now.
		cookie.Expires = time.Unix(1, 0)
	}

	// if SESSION_EXPIRE_ON_CLOSE = true
	if config.GetBool("SESSION_EXPIRE_ON_CLOSE") {
		cookie.Expires = GetMaxLifetime(options.Lifetime).Add(-10 * time.Minute)
		cookie.MaxAge = -1
	}

	return cookie
}
