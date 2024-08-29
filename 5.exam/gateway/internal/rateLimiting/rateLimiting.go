package ratelimiting

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	RateLimit int
	Window    time.Duration
}

func NewRateLimiter(ratelimiter int, window time.Duration)*RateLimiter{
	return &RateLimiter{RateLimit: ratelimiter, Window: window}
}

func (rl *RateLimiter) Limit(next http.HandlerFunc) http.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(rl.Window), rl.RateLimit)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
