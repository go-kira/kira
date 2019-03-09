package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-kira/kira"
	"github.com/go-kira/kira/tests/app/middleware"
)

func TestMiddleware(t *testing.T) {
	app := kira.New()
	app.Use(middleware.New())

	app.Get("/middleware", func(c *kira.Context) {
		c.String(" content ")
	})

	// Server
	s := httptest.NewServer(app.RegisterRoutes())

	// Request
	res, _ := http.Get(url(s, "/middleware"))
	defer res.Body.Close()

	// Get the response as string.
	content := contentS(res.Body)

	// Assert
	if content != "Before content After" {
		t.Errorf("expect: `Before content After`, have: %s", content)
	}
}

func TestMoreThanMiddleware(t *testing.T) {
	app := kira.New()
	app.Use(middleware.New())
	app.Use(middleware.New2())

	app.Get("/middleware", func(c *kira.Context) {
		c.String(" content ")
	})

	// Server
	s := httptest.NewServer(app.RegisterRoutes())

	// Request
	res, _ := http.Get(url(s, "/middleware"))
	defer res.Body.Close()

	// Get the response as string.
	content := contentS(res.Body)

	// Assert
	if content != "Before2Before content AfterAfter2" {
		t.Errorf("expect: `Before2Before content AfterAfter2`, have: %s", content)
	}
}

func TestRouteMiddleware(t *testing.T) {
	app := kira.New()

	app.Get("/route_middleware", func(c *kira.Context) {
		c.String(" inside ")
	}).Use(middleware.New())

	// Server
	s := httptest.NewServer(app.RegisterRoutes())

	// Request
	res, _ := http.Get(url(s, "/route_middleware"))
	defer res.Body.Close()

	// Get the response as string.
	content := contentS(res.Body)

	// Assert
	if content != "Before inside After" {
		t.Errorf("expect: `Before inside After`, have: %s", content)
	}
}

func TestMoreThanRouteMiddleware(t *testing.T) {
	app := kira.New()

	app.Get("/route_middleware", func(c *kira.Context) {
		c.String(" inside ")
	}).Use(middleware.New(), middleware.New2())

	// Server
	s := httptest.NewServer(app.RegisterRoutes())

	// Request
	res, _ := http.Get(url(s, "/route_middleware"))
	defer res.Body.Close()

	// Get the response as string.
	content := contentS(res.Body)

	// Assert
	if content != "Before2Before inside AfterAfter2" {
		t.Errorf("expect: `Before2Before inside AfterAfter2`, have: %s", content)
	}
}
