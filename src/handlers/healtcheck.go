package handlers

import "net/http"

// HealthCheckHandler handles the HTTP request for the health check endpoint.
// It responds with a 200 OK status and writes "OK" to indicate that the service is healthy.
//
// Parameters:
//   - w (http.ResponseWriter): The response writer to send the HTTP response.
//   - r (*http.Request): The HTTP request.
//
// This handler responds with a plain text "OK" message.

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
