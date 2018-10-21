package csrf

import (
	"github.com/Lafriakh/kira/helpers"
	"github.com/go-kira/kon"
	"golang.org/x/net/xsrftoken"
)

// Token - Get the CSRF token value.
func (c *CSRF) Token() string {
	return helpers.EncodeURLBase64([]byte(c.App.Session.Get("_token").(string)))
}

// RegenerateToken - Regenerate the CSRF token value.
func (c *CSRF) RegenerateToken(config *kon.Kon) {
	csrf := xsrftoken.Generate(config.GetString("APP_KEY"), "", "")

	// add the token to the session
	c.App.Session.Put("_token", csrf)
}
