package rate_limiter

import (
	"sync"
	"time"
)

const (
	MEMORY = 1
	REDIS  = 2
)

var (
	limiter RateLimiterInterface
	once    sync.Once
)

type RateLimiterInterface interface {
	AllowIP(ip string, maxRequests int, window time.Duration) bool
	AllowToken(token string, maxRequests int, window time.Duration) bool
	Cleanup()
}
