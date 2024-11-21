package handlers

import (
	"net/http"
)

// ExperimentalHandler handles the HTTP request for the experimental endpoint.
// It responds with a 200 OK status and writes a simple message indicating that
// it is an experimental endpoint.
//
// Parameters:
//   - w (http.ResponseWriter): The response writer to send the HTTP response.
//   - r (*http.Request): The HTTP request.
//
// This handler responds with a plain text "Experimental Endpoint" message.

func ExperimentalHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Experimental Endpoint"))
}
