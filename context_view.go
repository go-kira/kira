package kira

import (
	"bytes"
	"os"
	"strings"
)

// View send an html/template with an HTTP reply.
func (c *Context) View(temps string, data ...interface{}) error {
	// Set content type
	c.Response().Header().Set("Content-Type", "text/html")

	// parse the tempaltes
	template, templateData, err := parseView(c, temps, data)
	if err != nil {
		return err
	}

	// execute the templates
	err = template.Execute(c.Response(), templateData)
	if err != nil {
		return err
	}

	return nil
}

func (c *Context) ViewToString(temps string, data ...interface{}) (string, error) {
	template, templateData, err := parseView(c, temps, data)
	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	err = template.Execute(buf, templateData)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Validate if the view exists.
func (c *Context) ViewExists(tmp string) bool {
	fileSuffix := c.Config().GetString("views.file_suffix", ".go.html")
	viewPath := c.Config().GetString("views.path", "app/views/")

	templatePath := strings.Join([]string{viewPath, tmp, fileSuffix}, "")
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return false
	}

	return true
}
