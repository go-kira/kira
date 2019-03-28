package tests

import (
	"net/http"
	"testing"

	"github.com/go-kira/kira"
)

func TestView(t *testing.T) {
	s := endpoint("GET", "/method", func(c *kira.Context) {
		c.View("hello")
	})

	// Request
	res, _ := http.Get(url(s, "/method"))
	defer res.Body.Close()

	// Get the response as string.
	content := contentS(res.Body)

	// Assert
	if content != "Hello Kira!" {
		t.Errorf("expect: `Hello Kira!`, have: %s", content)
	}
}

func TestViewTemplate(t *testing.T) {
	s := endpoint("GET", "/method", func(c *kira.Context) {
		c.View("base|template")
	})

	// Request
	res, _ := http.Get(url(s, "/method"))
	defer res.Body.Close()

	// Get the response as string.
	content := contentS(res.Body)

	// Assert
	if content != "Base Template" {
		t.Errorf("expect: `Base Template`, have: %s", content)
	}
}

func TestViewTemplateInclude(t *testing.T) {
	s := endpoint("GET", "/method", func(c *kira.Context) {
		c.View("include")
	})

	// Request
	res, _ := http.Get(url(s, "/method"))
	defer res.Body.Close()

	// Get the response as string.
	content := contentS(res.Body)

	// Assert
	if content != "Included" {
		t.Errorf("expect: `Included`, have: %s", content)
	}
}

func TestViewString(t *testing.T) {
	s := endpoint("GET", "/method", func(c *kira.Context) {
		st, err := c.ViewToString("hello")
		if err != nil {
			c.Error(err)
		}

		c.WriteString(st)
	})

	// Request
	res, _ := http.Get(url(s, "/method"))
	defer res.Body.Close()

	// Get the response as string.
	content := contentS(res.Body)

	// Assert
	if content != "Hello Kira!" {
		t.Errorf("expect: `Hello Kira!`, have: %s", content)
	}
}
