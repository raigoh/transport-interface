package middleware

import (
	"bytes"
	"net/http"
	"time"
)

var cache = make(map[string]cachedResponse)

type cachedResponse struct {
	response    []byte
	contentType string
	expiration  time.Time
}

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
