package rate_limiter

import (
	"github.com/leonardopinho/GoLang/2.Rate_Limiter/config"
	"sync"
	"time"
)

var (
	limiter *MemoryRateLimiter
	once    sync.Once
)

type MemoryRateLimiter struct {
	config    config.Config
	mu        sync.Mutex
	ipData    map[string]*RateLimitConfig
	tokenData map[string]*RateLimitConfig
}

func NewInMemoryRateLimiter(cfg config.Config) *MemoryRateLimiter {
	once.Do(func() {
		limiter = &MemoryRateLimiter{
			config:    cfg,
			ipData:    make(map[string]*RateLimitConfig),
			tokenData: make(map[string]*RateLimitConfig),
		}
	})
	limiter.Cleanup()
	return limiter
}

func (r *MemoryRateLimiter) AllowIP(ip string, maxRequests int, window time.Duration) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	cfg, exists := r.ipData[ip]
	if !exists {
		cfg = &RateLimitConfig{
			Requests:  maxRequests,
			Window:    window,
			LastReset: time.Now(),
			Count:     0,
		}
		r.ipData[ip] = cfg
	}

	return checkAndIncrement(cfg, maxRequests, window)
}

func (r *MemoryRateLimiter) AllowToken(token string, maxRequests int, window time.Duration) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	cfg, exists := r.tokenData[token]
	if !exists {
		cfg = &RateLimitConfig{
			Requests:  maxRequests,
			Window:    window,
			LastReset: time.Now(),
			Count:     0,
		}
		r.tokenData[token] = cfg
	}

	return checkAndIncrement(cfg, maxRequests, window)
}

func (r *MemoryRateLimiter) Cleanup() {
	go func() {
		for {
			r.mu.Lock()
			now := time.Now()

			for key, rateCfg := range r.ipData {
				if now.Sub(rateCfg.LastReset) > time.Duration(r.config.BlockTimeRateLimit)*time.Second {
					delete(r.ipData, key)
				}
			}

			for key, rateCfg := range r.tokenData {
				if now.Sub(rateCfg.LastReset) > time.Duration(r.config.BlockTimeRateLimit)*time.Second {
					delete(r.tokenData, key)
				}
			}

			r.mu.Unlock()
		}
	}()
}

func checkAndIncrement(cfg *RateLimitConfig, maxRequests int, window time.Duration) bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	now := time.Now()
	if now.Sub(cfg.LastReset) >= window {
		cfg.Count = 0
		cfg.LastReset = now
	}

	if cfg.Count >= maxRequests {
		return false
	}

	cfg.Count++

	return true
}
