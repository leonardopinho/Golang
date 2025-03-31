package rate_limiter

import (
	"sync"
	"time"
)

const (
	IP = iota
	TOKEN
)

type RateLimitConfig struct {
	Requests    int
	RequestType int
	Window      time.Duration
	LastReset   time.Time
	Count       int
	mu          sync.Mutex
}
