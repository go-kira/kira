package kira

import (
	"github.com/Lafriakh/kira/session"
	"github.com/go-kira/kon"
)

// setupSession - for setup the framework session.
func setupSession(config *kon.Kon) *session.Session {
	name := config.GetString("SESSION_COOKIE")

	switch config.GetString("SESSION_DRIVER") {
	case "file":
		return session.NewSession(config, sessionFileHandler(config), name)
	}

	return session.NewSession(config, sessionFileHandler(config), name)
}

// sessionFileHandler - return a file handler for the session.
func sessionFileHandler(config *kon.Kon) session.Handler {
	path := config.GetString("SESSION_FILES")
	lifetime := config.GetInt("SESSION_COOKIE_LIFETIME")

	return session.NewFileHandler(path, lifetime)
}
