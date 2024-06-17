package handlers

import (
	"log"
	"net/http"
)

// Middleware to recover from panics and return a friendly error message
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
