package handlers

import (
	"net/http"
	"path/filepath"
)

// RenderErrorPage renders a custom error page with a 500 Internal Server Error status.
// It sets the HTTP response status to 500, then renders the error.html template.
//
// Parameters:
//   - w (http.ResponseWriter): The ResponseWriter used to write the HTTP response.
//
// This function is typically used to display a generic error page to the user
// when an internal server error occurs.

func RenderErrorPage(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	RenderTemplate(w, filepath.Join("templates", "error.html"), nil)
}
