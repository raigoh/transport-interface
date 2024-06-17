package middleware

import (
	"net/http"
	"sync"
	"time"
)

type rateLimiter struct {
	tokens         int
	maxTokens      int
	fillInterval   time.Duration
	lastTokenAdded time.Time
	mu             sync.Mutex
}

func newRateLimiter(tokens int, fillInterval time.Duration) *rateLimiter {
	return &rateLimiter{
		tokens:         tokens,
		maxTokens:      tokens,
		fillInterval:   fillInterval,
		lastTokenAdded: time.Now(),
	}
}

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

var visitors = make(map[string]*rateLimiter)
var mu sync.Mutex

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

func init() {
	go cleanupVisitors()
}

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
