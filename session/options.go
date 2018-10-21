package session

import "github.com/go-kira/kon"

// Options for session
type Options struct {
	Name          string
	Path          string
	Domain        string
	Secure        bool
	HTTPOnly      bool
	Lifetime      int
	ExpireOnClose bool
	FilesPath     string
}

func prepareOptions(config *kon.Kon) Options {
	var options Options

	// name
	if len(options.Name) == 0 {
		if config.IsNil("SESSION_COOKIE") {
			options.Name = "kira_session"
		} else {
			options.Name = config.GetString("SESSION_COOKIE")
		}
	}
	// path
	if len(options.Path) == 0 {
		if config.IsNil("SESSION_COOKIE_PATH") {
			options.Path = "/"
		} else {
			options.Path = config.GetString("SESSION_COOKIE_PATH")
		}

	}
	// domain
	if len(options.Domain) == 0 {
		if config.IsNil("SESSION_COOKIE_DOMAIN") {
			options.Domain = ""
		} else {
			options.Domain = config.GetString("SESSION_COOKIE_DOMAIN")
		}
	}
	// secure
	if !options.Secure {
		options.Secure = config.GetBool("SESSION_COOKIE_SECURE")
	}
	// http only
	if !options.HTTPOnly {
		options.HTTPOnly = config.GetBool("SESSION_COOKIE_HTTP_ONLY")
	}
	// lifetime
	if options.Lifetime == 0 {
		if config.IsNil("SESSION_LIFETIME") {
			options.Lifetime = 3600
		} else {
			options.Lifetime = config.GetInt("SESSION_LIFETIME")
		}
	}
	// expired on close
	if !options.ExpireOnClose {
		options.ExpireOnClose = config.GetBool("SESSION_EXPIRE_ON_CLOSE")
	}
	// files path
	if len(options.FilesPath) == 0 {
		if config.IsNil("SESSION_FILES") {
			options.FilesPath = "storage/framework/sessions"
		} else {
			options.FilesPath = config.GetString("SESSION_FILES")
		}
	}

	return options
}
