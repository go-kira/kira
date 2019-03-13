package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-kira/kira"
)

func TestNotFound(t *testing.T) {
	s := endpoint("GET", "/method", func(c *kira.Context) {
		c.String("GET")
	})

	// Request
	res, _ := http.Get(url(s, "/not_exists"))
	defer res.Body.Close()

	// Assert
	if res.StatusCode != 404 {
		t.Errorf("expect stauts: `404`, have: %d", res.StatusCode)
	}
}

func TestCustomNotFound(t *testing.T) {
	app := kira.New()

	app.NotFound(func(c *kira.Context) {
		c.Status(http.StatusNotFound)
		c.String("custom 404")
	})

	// Server
	s := httptest.NewServer(app.RegisterRoutes())

	// Request
	res, _ := http.Get(url(s, "/not_exists"))
	defer res.Body.Close()

	// Assert
	if res.StatusCode != 404 {
		t.Errorf("expect stauts: `404`, have: %d", res.StatusCode)
	}
	// Assert
	// Get the response as string.
	content := contentS(res.Body)

	// Assert
	if content != "custom 404" {
		t.Errorf("expect: `custom 404`, have: %s", content)
	}
}
