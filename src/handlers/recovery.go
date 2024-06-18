package handlers

import (
	"log"
	"net/http"
)

// WithRecovery is a middleware function that wraps an HTTP handler function
// to recover from panics that occur during its execution. If a panic occurs,
// it logs the panic message and then renders an error page to provide a
// friendly error message to the client.
//
// Parameters:
//   - handler (http.HandlerFunc): The HTTP handler function to wrap with recovery.
//
// Returns:
//   - http.HandlerFunc: A new HTTP handler function that includes recovery from panics.
//
// This middleware ensures that if a panic occurs during the execution of SomeHandler,
// the panic is logged and an error page is rendered instead of crashing the server.

func WithRecovery(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Panic occurred: %v", rec)
				RenderErrorPage(w)
			}
		}()
		handler(w, r)
	}
}
