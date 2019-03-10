package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-kira/kira"
)

func TestRouterGroup(t *testing.T) {
	app := kira.New()
	mfs := &mockFileSystem{}

	app.Group("prefix", func(g kira.Group) {
		g.Get("/get", func(c *kira.Context) {
			c.String("get")
		})
		g.Head("/head", func(c *kira.Context) {
			// The body should be empty
		})
		g.Post("/post", func(c *kira.Context) {
			c.String("post")
		})
		g.Put("/put", func(c *kira.Context) {
			c.String("put")
		})
		g.Patch("/patch", func(c *kira.Context) {
			c.String("patch")
		})
		g.Delete("/delete", func(c *kira.Context) {
			c.String("delete")
		})
		g.Options("/options", func(c *kira.Context) {
			// The body should be empty
		})

		g.ServeFiles("/files/*filepath", mfs)
	})

	// Server
	s := httptest.NewServer(app.RegisterRoutes())

	// === GET ===
	res, _ := http.Get(url(s, "/prefix/get"))
	defer res.Body.Close()
	content := contentS(res.Body)

	// Assert
	if res.StatusCode != http.StatusOK {
		t.Errorf("expect stauts: `202`, have: %d", res.StatusCode)
	}
	if content != "get" {
		t.Errorf("expect: `get`, have: %s", content)
	}

	// === HEAD ===
	resHEAD, _ := http.Head(url(s, "/prefix/head"))
	defer resHEAD.Body.Close()
	content = contentS(resHEAD.Body)

	// Assert
	if resHEAD.StatusCode != http.StatusOK {
		t.Errorf("expect stauts: `202`, have: %d", resHEAD.StatusCode)
	}

	// === POST ===
	resPOST, _ := http.Post(url(s, "/prefix/post"), "application/json", nil)
	defer resPOST.Body.Close()
	content = contentS(resPOST.Body)

	// Assert
	if resPOST.StatusCode != http.StatusOK {
		t.Errorf("expect stauts: `202`, have: %d", resPOST.StatusCode)
	}
	if content != "post" {
		t.Errorf("expect: `post`, have: %s", content)
	}

	// === Put ===
	reqPut, _ := http.NewRequest(http.MethodPut, url(s, "/prefix/put"), nil)
	reqPut.Header.Set("Content-Type", "application/json; charset=utf-8")
	resPut, _ := http.DefaultClient.Do(reqPut)
	defer resPut.Body.Close()
	content = contentS(resPut.Body)

	// Assert
	if resPut.StatusCode != http.StatusOK {
		t.Errorf("expect stauts: `202`, have: %d", resPut.StatusCode)
	}
	if content != "put" {
		t.Errorf("expect: `put`, have: %s", content)
	}

	// === Patch ===
	reqPatch, _ := http.NewRequest(http.MethodPatch, url(s, "/prefix/patch"), nil)
	reqPatch.Header.Set("Content-Type", "application/json; charset=utf-8")
	resPatch, _ := http.DefaultClient.Do(reqPatch)
	defer resPatch.Body.Close()
	content = contentS(resPatch.Body)

	// Assert
	if resPatch.StatusCode != http.StatusOK {
		t.Errorf("expect stauts: `202`, have: %d", resPatch.StatusCode)
	}
	if content != "patch" {
		t.Errorf("expect: `patch`, have: %s", content)
	}

	// === Delete ===
	reqDelete, _ := http.NewRequest(http.MethodDelete, url(s, "/prefix/delete"), nil)
	reqDelete.Header.Set("Content-Type", "application/json; charset=utf-8")
	resDelete, _ := http.DefaultClient.Do(reqDelete)
	defer resDelete.Body.Close()
	content = contentS(resDelete.Body)

	// Assert
	if resDelete.StatusCode != http.StatusOK {
		t.Errorf("expect stauts: `202`, have: %d", resDelete.StatusCode)
	}
	if content != "delete" {
		t.Errorf("expect: `delete`, have: %s", content)
	}

	// === Options ===
	reqOptions, _ := http.NewRequest(http.MethodOptions, url(s, "/prefix/options"), nil)
	reqOptions.Header.Set("Content-Type", "application/json; charset=utf-8")
	resOptions, _ := http.DefaultClient.Do(reqOptions)
	defer resOptions.Body.Close()
	content = contentS(resOptions.Body)

	// Assert
	if resOptions.StatusCode != http.StatusOK {
		t.Errorf("expect stauts: `202`, have: %d", resOptions.StatusCode)
	}

	// === ServeFiles ===
	resServFiles, _ := http.Get(url(s, "/prefix/files/favicon.ico"))
	defer resServFiles.Body.Close()

	// Assert
	if resServFiles.StatusCode == http.StatusOK {
		t.Errorf("expect stauts: `202`, have: %d", resServFiles.StatusCode)
	}
	if !mfs.opened {
		t.Error("serving file failed")
	}
}
