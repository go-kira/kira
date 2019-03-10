package tests

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http/httptest"

	"github.com/go-kira/kira"
)

func endpoint(method string, path string, handler kira.HandlerFunc) *httptest.Server {
	app := kira.New()

	switch method {
	case "GET":
		app.Get(path, handler)
	case "HEAD":
		app.Head(path, handler)
	case "POST":
		app.Post(path, handler)
	case "PUT":
		app.Put(path, handler)
	case "PATCH":
		app.Patch(path, handler)
	case "DELETE":
		app.Delete(path, handler)
	case "OPTIONS":
		app.Options(path, handler)
	}

	return httptest.NewServer(app.RegisterRoutes())
}

func url(server *httptest.Server, path string) string {
	return fmt.Sprintf("%s%s", server.URL, path)
}

func contentS(r io.Reader) string {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s", content)
}

func content(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}
