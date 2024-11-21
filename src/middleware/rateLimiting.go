package middleware

import (
	"net/http"
	"sync"
	"time"
)

// rateLimiter represents a token bucket rate limiter.
type rateLimiter struct {
	tokens         int           // Current number of tokens in the bucket
	maxTokens      int           // Maximum number of tokens the bucket can hold
	fillInterval   time.Duration // Interval at which tokens are added to the bucket
	lastTokenAdded time.Time     // Time when the last token was added to the bucket
	mu             sync.Mutex    // Mutex to synchronize access to the rate limiter
}

// newRateLimiter creates and initializes a new rate limiter.
func newRateLimiter(tokens int, fillInterval time.Duration) *rateLimiter {
	return &rateLimiter{
		tokens:         tokens,
		maxTokens:      tokens,
		fillInterval:   fillInterval,
		lastTokenAdded: time.Now(),
	}
}

// allow checks if a request can proceed based on the rate limiter's token availability.
// It returns true if the request is allowed (token available), false otherwise.
func (rl *rateLimiter) allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastTokenAdded)
	tokensToAdd := int(elapsed / rl.fillInterval)
	if tokensToAdd > 0 {
		rl.tokens += tokensToAdd
		if rl.tokens > rl.maxTokens {
			rl.tokens = rl.maxTokens
		}
		rl.lastTokenAdded = now
	}

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}

	return false
}

// visitors stores rate limiters for each IP address.
var visitors = make(map[string]*rateLimiter)
var mu sync.Mutex

// getVisitor retrieves or creates a rate limiter for the given IP address.
func getVisitor(ip string) *rateLimiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		v = newRateLimiter(3, time.Second) // 3 tokens, refill one token per second
		visitors[ip] = v
	}

	return v
}

// cleanupVisitors periodically cleans up expired rate limiters from the visitors map.
func cleanupVisitors() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastTokenAdded) > 5*time.Minute {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}

// init initializes the rate limiting middleware by starting the cleanup routine.
func init() {
	go cleanupVisitors()
}

// RateLimitingMiddleware is a middleware function that enforces rate limiting based on IP address.
// It restricts the number of requests a client can make within a certain time window.
//
// Parameters:
//   - next (http.Handler): The next HTTP handler in the chain.
//
// Returns:
//   - http.Handler: A new HTTP handler that includes rate limiting logic.
//
// This middleware helps protect web servers from abuse by limiting the rate of requests from individual IP addresses.
func RateLimitingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := getVisitor(r.RemoteAddr)
		if !limiter.allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
