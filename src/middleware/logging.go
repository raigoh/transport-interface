package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware is a middleware function that logs details of incoming HTTP requests.
// It logs the HTTP method, request URI, remote address, and request processing time.
//
// Parameters:
//   - next (http.Handler): The next HTTP handler in the chain.
//
// Returns:
//   - http.Handler: A new HTTP handler that includes logging logic.
//
// This middleware provides logging capabilities for HTTP requests, helping to monitor
// incoming traffic and track request performance.

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s %v", r.Method, r.RequestURI, r.RemoteAddr, time.Since(start))
	})
}
