package tests

import (
	"net/http"
	"testing"

	"github.com/go-kira/kira"
)

func TestGET(t *testing.T) {
	s := endpoint("GET", "/method", func(c *kira.Context) {
		c.WriteString("GET")
	})

	// Request
	res, _ := http.Get(url(s, "/method"))
	defer res.Body.Close()

	// Get the response as string.
	content := contentS(res.Body)

	// Assert
	if content != "GET" {
		t.Errorf("expect: `GET`, have: %s", content)
	}
}

func TestHEAD(t *testing.T) {
	s := endpoint("HEAD", "/method", func(c *kira.Context) {
		// The body should be empty
	})

	// Request
	res, _ := http.Head(url(s, "/method"))
	defer res.Body.Close()

	// Assert
	if res.StatusCode != http.StatusOK {
		t.Errorf("expect status: `200`, have: %d", res.StatusCode)
	}
}

func TestPOST(t *testing.T) {
	s := endpoint("POST", "/method", func(c *kira.Context) {
		c.WriteString("POST")
	})

	// Request
	res, _ := http.Post(url(s, "/method"), "application/json", nil)
	defer res.Body.Close()

	// Get the response as string.
	content := contentS(res.Body)

	// Assert
	if content != "POST" && res.StatusCode == http.StatusOK {
		t.Errorf("expect: `POST`, have: %s", content)
	}
}

func TestPUT(t *testing.T) {
	s := endpoint("PUT", "/method", func(c *kira.Context) {
		c.WriteString("PUT")
	})

	// Request
	req, _ := http.NewRequest(http.MethodPut, url(s, "/method"), nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, _ := http.DefaultClient.Do(req)

	// Get the response as string.
	content := contentS(res.Body)

	// Assert
	if content != "PUT" && res.StatusCode == http.StatusOK {
		t.Errorf("expect: `PUT`, have: %s", content)
	}
}

func TestPATCH(t *testing.T) {
	s := endpoint("PATCH", "/method", func(c *kira.Context) {
		c.WriteString("PATCH")
	})

	// Request
	req, _ := http.NewRequest(http.MethodPatch, url(s, "/method"), nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, _ := http.DefaultClient.Do(req)

	// Get the response as string.
	content := contentS(res.Body)

	// Assert
	if content != "PATCH" && res.StatusCode == http.StatusOK {
		t.Errorf("expect: `PATCH`, have: %s", content)
	}
}

func TestDELETE(t *testing.T) {
	s := endpoint("DELETE", "/method", func(c *kira.Context) {
		c.WriteString("DELETE")
	})

	// Request
	req, _ := http.NewRequest(http.MethodDelete, url(s, "/method"), nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, _ := http.DefaultClient.Do(req)

	// Get the response as string.
	content := contentS(res.Body)

	// Assert
	if content != "DELETE" && res.StatusCode == http.StatusOK {
		t.Errorf("expect: `DELETE`, have: %s", content)
	}
}

func TestOPTIONS(t *testing.T) {
	s := endpoint("OPTIONS", "/method", func(c *kira.Context) {
		// No body for this method.
	})

	// Request
	req, _ := http.NewRequest(http.MethodOptions, url(s, "/method"), nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, _ := http.DefaultClient.Do(req)

	// Assert
	if res.StatusCode != http.StatusOK {
		t.Errorf("expect status: `200`, have: %d", res.StatusCode)
	}
}
