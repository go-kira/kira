package kira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Lafriakh/log"
)

// View - struct for views logic.
type View struct {
	App        *App
	Data       map[string]interface{}
	customPath string
}

// Data - type of view data.
type Data map[string]interface{}

// Data - to put data into views.
func (a *App) Data(data Data) {
	a.View.Data = data
}

// Render - short access to the view render method.
func (a *App) Render(w http.ResponseWriter, req *http.Request, templates ...string) {
	a.View.Render(w, req, templates...)
}

// AddPath to add a custom path to the view struct.
func (v *View) AddPath(path string) error {
	// check if this path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	v.customPath = path

	return nil
}

// Render template
func (v *View) Render(w http.ResponseWriter, req *http.Request, templates ...string) {
	// this will be used to catch the error
	buf := &bytes.Buffer{}

	// hold all templates
	var templatesFiles []string
	baseTemplate := filepath.Base(templates[0]) + v.App.Configs.GetString("VIEWS_FILE_SUFFIX")
	// loop throw all templates
	for _, temp := range templates {
		if _, err := os.Stat(filepath.Join(v.customPath, temp+v.App.Configs.GetString("VIEWS_FILE_SUFFIX"))); err == nil {
			templatesFiles = append(templatesFiles, filepath.Join(v.customPath, temp+v.App.Configs.GetString("VIEWS_FILE_SUFFIX")))
		} else {
			templatesFiles = append(templatesFiles, v.App.Configs.GetString("VIEWS_PATH")+temp+v.App.Configs.GetString("VIEWS_FILE_SUFFIX"))
		}
	}

	// parse templates
	Template, err := template.New(baseTemplate).Funcs(v.FuncMap(req)).ParseFiles(templatesFiles...)
	if err != nil {
		log.Panic(err)
	}

	// global variables
	v.Global(req)
	err = Template.Execute(buf, v.Data)
	// reset view data before any error
	v.Data = map[string]interface{}{}

	// check for errors
	if err != nil {
		v.App.Log.Panic(err)
	}

	// write the response
	buf.WriteTo(w)
}

// Global ...
func (v *View) Global(req *http.Request) {
	errors := &ViewErrors{}
	errors.Errors = v.App.Session.GetWithDefault("errors", map[string]string{}).(map[string]string)

	// site
	v.Data["site_name"] = v.App.Configs.GetString("SITE_NAME")
	// request id
	v.Data["request_id"] = req.Context().Value(v.App.Configs.GetString("SERVER_REQUEST_ID"))
	// errors
	v.Data["errors"] = errors
	// csrf
	v.Data["csrf"] = req.Context().Value("csrf")
	v.Data["csrfField"] = template.HTML(fmt.Sprintf(`<input type="hidden" name="%s" value="%s">`, v.App.Configs.GetString("CSRF_FIELD_NAME"), req.Context().Value("csrf")))
}

// FuncMap - return a collection of teplate functions.
func (v *View) FuncMap(req *http.Request) template.FuncMap {
	return template.FuncMap{
		"session": func(key interface{}) interface{} {
			return v.App.Session.Get(key)
		},
		"config": func(key string) interface{} {
			return v.App.Configs.Get(key)
		},
		"include": func(filename string) interface{} {
			// read the template content.
			b, err := ioutil.ReadFile(v.App.Configs.GetString("VIEWS_PATH") + filename + v.App.Configs.GetString("VIEWS_FILE_SUFFIX"))
			if err != nil {
				return nil
			}
			// create a buffer.
			var buffer bytes.Buffer
			// parse the template.
			tmpl, err := template.New("").Funcs(v.FuncMap(req)).Parse(string(b))
			if err != nil {
				log.Panic(err)
			}
			// add global variables to the included template
			v.Global(req)

			// execute the template with the data.
			errs := tmpl.Execute(&buffer, v.Data)
			// check for errors
			if errs != nil {
				log.Panic(errs)
			}

			// return the template to the parent template.
			return template.HTML(buffer.String())
		},
		"url": func() string {
			return req.URL.Path
		},
		"join": func(s ...string) string {
			// first arg is sep, remaining args are strings to join
			return strings.Join(s[1:], s[0])
		},
	}
}

// JSON response.
func (v *View) JSON(w http.ResponseWriter, data interface{}) {
	// parse data to json format
	response, err := json.Marshal(data)
	if err != nil {
		log.Panic(err)
	}

	// return json with headers...
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// ViewErrors - to manage view errors.
type ViewErrors struct {
	Errors map[string]string
}

// New - append errors to the struct.
func (e *ViewErrors) New(data map[string]string) {
	e.Errors = data
}

// All - return all errors.
func (e *ViewErrors) All() map[string]string {
	return e.Errors
}

// Has - check if the key exists, if exists return it's value.
func (e *ViewErrors) Has(key string) string {
	if _, ok := e.Errors[key]; ok {
		return e.Errors[key]
	}
	return ""
}
