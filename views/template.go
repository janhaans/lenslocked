package views

import (
	"fmt"
	"html/template"
	"net/http"
)

type Template struct {
	Template *template.Template
}

func (t *Template) Execute(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.Template.Execute(w, data)
	if err != nil {
		err = fmt.Errorf("HTML Template Execution failed %v", err)
		return err
	}
	return nil
}

func ParseTemplate(filepath string) (*Template, error) {
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		err = fmt.Errorf("HTML Template Parsing failed %v", err)
		return &Template{}, err
	}
	return &Template{tmpl}, nil
}
