package limiter

import "golang.org/x/time/rate"

// RateLimiter interface represents the rate limiter behavior.
type RateLimiterInterface interface {
	Allow() bool
	//bisa ditambahkan yang lain nantinya
}

// RealRateLimiter implements RateLimiter using the real rate limiter logic.
type RealRateLimiter struct {
	limiter *rate.Limiter
}

func NewRateLimiter(limit rate.Limit, burst int) *RealRateLimiter {
	return &RealRateLimiter{
		limiter: rate.NewLimiter(limit, burst),
	}
}

func (rl *RealRateLimiter) Allow() bool {
	return rl.limiter.Allow()
}
