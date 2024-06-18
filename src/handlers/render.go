package handlers

import (
	"html/template"
	"net/http"
)

// RenderTemplate renders an HTML template with data to the HTTP response.
// It parses the template located at tmplPath, executes it with the provided data,
// and writes the rendered HTML to the ResponseWriter.
//
// Parameters:
//   - w (http.ResponseWriter): The ResponseWriter used to write the HTTP response.
//   - tmplPath (string): The file path of the HTML template to render.
//   - data (interface{}): The data to pass into the template for rendering.
//
// This function handles parsing and executing an HTML template,
// responding with an appropriate HTTP error if any parsing or execution errors occur.

func RenderTemplate(w http.ResponseWriter, tmplPath string, data interface{}) {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Could not parse template", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}
