package rate_limiter

import "time"

type RateLimiterInterface interface {
	Allow(ip string, maxRequests int, window time.Duration) bool
	Cleanup() error
}
