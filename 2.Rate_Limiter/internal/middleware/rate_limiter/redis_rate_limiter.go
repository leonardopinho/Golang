package rate_limiter

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/leonardopinho/GoLang/2.Rate_Limiter/config"
	"sync"
	"time"
)

type RedisRateLimiter struct {
	client *redis.Client
	config config.Config
	mu     *sync.Mutex
}

func NewRedisRateLimiter(cfg config.Config) RateLimiterInterface {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	once.Do(func() {
		limiter = &RedisRateLimiter{
			client: rdb,
			config: cfg,
			mu:     &sync.Mutex{},
		}
	})

	limiter.Cleanup()
	return limiter
}

func (r *RedisRateLimiter) AllowIP(ip string, maxRequests int, window time.Duration) bool {
	ctx := context.Background()

	key := "rate_limit:" + ip

	count, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return false
	}

	if count == 1 {
		// Define a janela de tempo (expire) para a chave
		r.client.Expire(ctx, key, window)
	}

	// Se a contagem excedeu o limite, bloqueia
	if count > int64(maxRequests) {
		return false
	}

	return true
}

func (r *RedisRateLimiter) AllowToken(token string, maxRequests int, window time.Duration) bool {
	ctx := context.Background()

	key := "rate_limit:token:" + token

	count, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return false
	}

	if count == 1 {
		r.client.Expire(ctx, key, window)
	}

	if count > int64(maxRequests) {
		return false
	}

	return true
}

func (r *RedisRateLimiter) Cleanup() {
	//
}
