package tests

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/go-kira/kira"
)

func TestContextJSON(t *testing.T) {
	s := endpoint("POST", "/json", func(c *kira.Context) {
		if c.WantsJSON() {
			c.JSON(struct {
				Message string `json:"message"`
			}{"json"})
		}
	})

	// Request
	reqPost, _ := http.NewRequest(http.MethodPost, url(s, "/json"), nil)
	reqPost.Header.Set("Content-Type", "application/json; charset=utf-8")
	// We want the content to be in json.
	reqPost.Header.Set("Accept", "application/json")
	// Send request.
	resPost, _ := http.DefaultClient.Do(reqPost)
	// Assert
	content, err := content(resPost.Body)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(content, []byte("{\"message\":\"json\"}\n")) {
		t.Errorf(`expect: "{"message":"json"}", have: %s`, content)
	}

	if resPost.StatusCode != http.StatusOK {
		t.Errorf("expect status: `200`, have: %d", resPost.StatusCode)
	}
}

func TestContextJSONParse(t *testing.T) {
	s := endpoint("POST", "/json", func(c *kira.Context) {
		if c.WantsJSON() {
			st := struct {
				Message string `json:"message"`
			}{}

			c.ParseJSON(&st)

			if st.Message != "json" {
				t.Errorf(`expect: "{"message":"json"}", have: %+v`, st)
				return
			}
			c.String("parsed")
			return
		}
		c.String("not parsed")
	})

	jsonStr := []byte("{\"message\":\"json\"}\n")

	// Request
	reqPost, _ := http.NewRequest(http.MethodPost, url(s, "/json"), bytes.NewBuffer(jsonStr))
	reqPost.Header.Set("Content-Type", "application/json; charset=utf-8")
	// We want the content to be in json.
	reqPost.Header.Set("Accept", "application/json")
	// Send request.
	resPost, _ := http.DefaultClient.Do(reqPost)
	// Assert
	content := contentS(resPost.Body)

	if content != "parsed" && resPost.StatusCode == http.StatusOK {
		t.Errorf(`expect: "{"message":"json"}", have: %s`, content)
	}
}
