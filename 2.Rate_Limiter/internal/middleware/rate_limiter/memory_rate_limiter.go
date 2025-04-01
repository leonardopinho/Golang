package rate_limiter

import (
	"github.com/leonardopinho/GoLang/2.Rate_Limiter/config"
	"sync"
	"time"
)

type MemoryRateLimiter struct {
	config config.Config
	mu     *sync.Mutex
	data   map[string]*RateLimitConfig
}

func NewInMemoryRateLimiter(cfg config.Config) RateLimiterInterface {
	once.Do(func() {
		limiter = &MemoryRateLimiter{
			config: cfg,
			data:   make(map[string]*RateLimitConfig),
			mu:     &sync.Mutex{},
		}
	})
	limiter.Cleanup()
	return limiter
}

func (r *MemoryRateLimiter) AllowIP(ip string, maxRequests int, window time.Duration) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	cfg, exists := r.data[ip]
	if !exists {
		cfg = &RateLimitConfig{
			Requests:    maxRequests,
			RequestType: IP,
			Window:      window,
			LastReset:   time.Now(),
			Count:       0,
		}
		r.data[ip] = cfg
	}

	return checkAndIncrement(cfg, maxRequests, window)
}

func (r *MemoryRateLimiter) AllowToken(token string, maxRequests int, window time.Duration) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	// TODO: validação do token

	cfg, exists := r.data[token]
	if !exists {
		cfg = &RateLimitConfig{
			Requests:    maxRequests,
			RequestType: TOKEN,
			Window:      window,
			LastReset:   time.Now(),
			Count:       0,
		}
		r.data[token] = cfg
	}

	return checkAndIncrement(cfg, maxRequests, window)
}

func (r *MemoryRateLimiter) Cleanup() {
	go func() {
		for {
			r.mu.Lock()
			now := time.Now()

			for key, rateCfg := range r.data {
				if now.Sub(rateCfg.LastReset) > time.Duration(checkTimeLimit(rateCfg, r))*time.Second {
					delete(r.data, key)
				}
			}

			r.mu.Unlock()
		}
	}()
}

func checkTimeLimit(rateCfg *RateLimitConfig, r *MemoryRateLimiter) int {
	if rateCfg.RequestType == IP {
		return r.config.BlockTimeRateLimit
	}
	return r.config.BlockTimeRateLimitToken
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
