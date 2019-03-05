package kira

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

// View send an html/template with an HTTP reply.
func (c *Context) View(templates ...string) error {
	buf := &bytes.Buffer{}

	// templates extention
	fileSuffix := c.Config().GetString("views.file_suffix", ".go.html")
	// tempaltes path
	viewPath := c.Config().GetString("views.path", "app/views/")

	// hold all templates
	var templatesFiles []string
	baseTemplate := filepath.Base(templates[0]) + fileSuffix

	// loop throw all templates
	for _, temp := range templates {
		templatesFiles = append(templatesFiles, viewPath+temp+fileSuffix)
	}

	// parse templates
	Template, err := template.New(baseTemplate).ParseFiles(templatesFiles...)
	if err != nil {
		return err
	}

	err = Template.Execute(buf, c.data)
	// check for errors
	if err != nil {
		return err
	}

	// write the response
	buf.WriteTo(c.Response())

	return nil
}

// Validate if the view exists.
func (c *Context) viewExists(tmp string) bool {
	fileSuffix := c.Config().GetString("views.file_suffix", ".go.html")
	viewPath := c.Config().GetString("views.path", "app/views/")

	templatePath := strings.Join([]string{viewPath, tmp, fileSuffix}, "")
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return false
	}

	return true
}
