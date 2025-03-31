package rate_limiter

import "time"

type RateLimiterInterface interface {
	AllowIP(ip string, maxRequests int, window time.Duration) bool
	AllowToken(token string, maxRequests int, window time.Duration) bool
	Cleanup() error
}
