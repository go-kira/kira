package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-kira/kira"
	"github.com/go-kira/kira/tests/app/middleware"
)

func TestContextData(t *testing.T) {
	s := endpoint("GET", "/data", func(c *kira.Context) {
		// Set the data to the context.
		c.SetData("key", "value")

		if c.HasData("key") != true {
			t.Errorf("expect: `true`, have: %t", c.HasData("key"))
		}

		if c.HasData("key_not_exists") != false {
			t.Errorf("expect: `false`, have: %t", c.HasData("key_not_exists"))
		}

		// Get the data from the context.
		if c.GetData("key").(string) != "value" {
			t.Errorf("expect: `value`, have: %s", c.GetData("key").(string))
		}
	})

	// Request
	res, _ := http.Get(url(s, "/data"))
	defer res.Body.Close()

	// Assert
	if res.StatusCode != http.StatusOK {
		t.Errorf("expect stauts: `202`, have: %d", res.StatusCode)
	}
}

func TestContextDataFromMiddleware(t *testing.T) {
	app := kira.New()
	app.Use(middleware.NewContextData())

	app.Get("/middleware", func(c *kira.Context) {
		// Get the data from the context.
		if c.GetData("foo").(string) != "bar" {
			t.Errorf("expect: `value`, have: %s", c.GetData("key").(string))
		}
	})

	// Server
	s := httptest.NewServer(app.RegisterRoutes())

	// Request
	res, _ := http.Get(url(s, "/middleware"))
	defer res.Body.Close()

	// Assert
	if res.StatusCode != http.StatusOK {
		t.Errorf("expect stauts: `202`, have: %d", res.StatusCode)
	}
}
