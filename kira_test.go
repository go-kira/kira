package kira

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Create a server with initial routes.
func setupServer() *httptest.Server {
	app := New()
	app.LoadConfig("./testdata/config/application.yaml")

	app.Get("/kira", func(c *Context) {
		c.String("Kira")
	})
	app.Get("/kira_template", func(c *Context) {
		c.View("hello")
	})
	app.Get("/kira_template_not_exists", func(c *Context) {
		if err := c.View("not_exists"); err != nil {
			c.String("not exists")
		} else {
			c.String("exists")
		}
	})

	ts := httptest.NewServer(app.NewRouter())

	return ts
}

// Test against GET requests.
func getResponse(path string) ([]byte, error) {
	ts := setupServer()
	defer ts.Close()

	res, err := http.Get(fmt.Sprintf("%s%s", ts.URL, path))
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	return content, nil
}

func TestRunServer(t *testing.T) {
	content, err := getResponse("/kira")
	if err != nil {
		t.Error(err)
	}

	// Assert
	if fmt.Sprintf("%s", content) != "Kira" {
		t.Errorf("The response is not equal to: `Kira`, is equal to: %s", content)
	}
}

func TestTemplate(t *testing.T) {
	content, err := getResponse("/kira_template")
	if err != nil {
		t.Error(err)
	}

	// Assert
	if fmt.Sprintf("%s", content) != "Kira Template" {
		t.Errorf("The response is not equal to: `Kira Template`, is equal to: %s", content)
	}
}
func TestTemplateNotExists(t *testing.T) {
	content, err := getResponse("/kira_template_not_exists")
	if err != nil {
		t.Error(err)
	}

	// Assert
	if fmt.Sprintf("%s", content) != "not exists" {
		t.Errorf("The response is not equal to: `not exists`, is equal to: %s", content)
	}
}
