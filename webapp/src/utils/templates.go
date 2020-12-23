package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates insert html templates on var templates
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecutingTemplate func
func ExecutingTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
