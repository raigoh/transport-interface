package handlers

import (
	"net/http"
	"path/filepath"
)

// Function to render a custom error page
func RenderErrorPage(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	RenderTemplate(w, filepath.Join("templates", "error.html"), nil)
}
