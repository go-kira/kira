package kira

import (
	"bytes"
	"html/template"
	"path/filepath"
)

// View send an html/template with an HTTP reply.
func (c *Context) View(templates ...string) error {
	buf := &bytes.Buffer{}

	// templates extention
	fileSuffix := c.Config().GetDefault("VIEWS_FILE_SUFFIX", ".go.html").(string)
	// tempaltes path
	viewPath := c.Config().GetDefault("VIEWS_PATH", "app/views/").(string)

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
