package middleware

import (
	"bytes"
	"net/http"
	"time"
)

var cache = make(map[string]cachedResponse)

// cachedResponse represents a cached HTTP response.
type cachedResponse struct {
	response    []byte    // response body as bytes
	contentType string    // content type of the response
	expiration  time.Time // expiration time of the cached response
}

// CachingMiddleware is a middleware function that implements caching for GET requests.
// It caches HTTP responses based on the request URI for a specified duration.
//
// Parameters:
//   - next (http.Handler): The next HTTP handler in the chain.
//   - duration (time.Duration): The duration for which the response should be cached.
//
// Returns:
//   - http.Handler: A new HTTP handler that includes caching logic.
//
// This middleware caches responses of GET requests and serves cached responses
// if available and not expired, reducing the load on backend servers.

func CachingMiddleware(next http.Handler, duration time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			if cached, found := cache[r.RequestURI]; found && time.Now().Before(cached.expiration) {
				w.Header().Set("Content-Type", cached.contentType)
				w.Write(cached.response)
				return
			}
		}

		rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rec, r)

		if r.Method == http.MethodGet && rec.statusCode == http.StatusOK {
			cache[r.RequestURI] = cachedResponse{
				response:    rec.body.Bytes(),
				contentType: rec.Header().Get("Content-Type"),
				expiration:  time.Now().Add(duration),
			}
		}
	})
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
