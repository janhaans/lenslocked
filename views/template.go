package views

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/janhaans/lenslocked/templates"
)

type Template struct {
	Template *template.Template
}

func (t *Template) Execute(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.Template.Execute(w, data)
	if err != nil {
		err = fmt.Errorf("HTML Template Execution failed: %v", err)
		return err
	}
	return nil
}

/*
func ParseTemplate(filepath string) (*Template, error) {
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		err = fmt.Errorf("HTML template parsing failed: %v", err)
		return &Template{}, err
	}
	return &Template{tmpl}, nil
}
*/

func ParseFS(filepath string) (*Template, error) {
	tmpl, err := template.ParseFS(templates.FS, filepath)
	if err != nil {
		err = fmt.Errorf("HTML template parsing failed: %v", err)
		return &Template{}, err
	}
	return &Template{tmpl}, nil
}

func Must(tmpl *Template, err error) *Template {
	if err != nil {
		panic(err)
	}
	return tmpl
}
