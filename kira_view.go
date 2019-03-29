package kira

import (
	"encoding/json"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
)

// ViewData represent the view data.
type ViewData map[string]interface{}

// parse the view and return the view template and the view data.
func parseView(c *Context, temps string, data ...interface{}) (*template.Template, error) {
	fileSuffix := c.Config().GetString("views.file_suffix", ".go.html")
	viewPath := c.Config().GetString("views.path", "./app/views/")

	templates := strings.Split(temps, "|")

	// hold all templates
	var templatesFiles []string
	baseTemplate := filepath.Base(templates[0]) + fileSuffix

	// loop throw all templates
	for _, temp := range templates {
		tmplPath := filepath.Join(viewPath, temp+fileSuffix)
		if !c.ViewExists(temp) {
			return nil, fmt.Errorf("kira: template %s not exits", tmplPath)
		}

		templatesFiles = append(templatesFiles, tmplPath)
	}

	// parse templates
	template, err := template.New(baseTemplate).Funcs(viewFuncs(c)).ParseFiles(templatesFiles...)
	if err != nil {
		return nil, err
	}

	return template, nil
}

func parseViewData(data ...interface{}) interface{} {
	if len(data) > 0 {
		return data[0]
	}
	return nil
}

// default views functions.
func viewFuncs(ctx *Context) template.FuncMap {
	return template.FuncMap{
		"env": func() string {
			return ctx.Env()
		},
		"data": func(key string) interface{} {
			return ctx.GetData(key)
		},
		"config": func(key string) interface{} {
			return ctx.Config().Get(key)
		},
		"url": func() string {
			return ctx.Request().URL.Path
		},
		"join": func(s ...string) string {
			// first arg is sep, remaining args are strings to join
			return strings.Join(s[1:], s[0])
		},
		"include": func(filename string) (string, error) {
			return ctx.ViewToString(filename)
		},
		"json": func(v interface{}) template.JS {
			js, _ := json.Marshal(v)
			return template.JS(js)
		},
	}
}
