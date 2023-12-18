package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // Set the content type to text/html
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Something went wrong executing", http.StatusInternalServerError)
		return
	}
}

func Parse(filepath string) (Template, error) {
	htmlTpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		return Template{}, fmt.Errorf("error parsing template: %w", err)
	}
	return Template{
		htmlTpl: htmlTpl,
	}, nil
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, pattern string) (Template, error) {
	htmlTpl, err := template.ParseFS(fs, pattern)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		return Template{}, fmt.Errorf("error parsing template: %w", err)
	}
	return Template{
		htmlTpl: htmlTpl,
	}, nil
}
