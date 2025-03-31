package rate_limiter

import (
	"sync"
	"time"
)

type RateLimitConfig struct {
	Requests  int
	Window    time.Duration
	LastReset time.Time
	Count     int
	mu        sync.Mutex
}
