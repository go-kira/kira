package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-kira/kira"
)

type mockFileSystem struct {
	opened bool
}

func (mfs *mockFileSystem) Open(name string) (http.File, error) {
	mfs.opened = true
	return nil, errors.New("this is just a mock")
}

func TestRouterServeFiles(t *testing.T) {
	mfs := &mockFileSystem{}

	app := kira.New()
	app.ServeFiles("/*filepath", mfs)

	// Server
	s := httptest.NewServer(app.RegisterRoutes())

	// Request
	res, _ := http.Get(url(s, "/favicon.ico"))
	defer res.Body.Close()

	// Assert
	if res.StatusCode == http.StatusOK {
		t.Errorf("expect stauts: `202`, have: %d", res.StatusCode)
	}

	if !mfs.opened {
		t.Error("serving file failed")
	}
}
