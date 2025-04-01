package middleware

import (
	"github.com/leonardopinho/GoLang/2.Rate_Limiter/config"
	"github.com/leonardopinho/GoLang/2.Rate_Limiter/internal/middleware/rate_limiter"
	"log"
	"net"
	"net/http"
	"runtime/debug"
	"time"
)

var (
	limiter rate_limiter.RateLimiterInterface
)

func RateLimiterMiddleware(cfg *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if cfg.Strategy == rate_limiter.MEMORY {
				limiter = rate_limiter.NewInMemoryRateLimiter(*cfg)
			} else if cfg.Strategy == rate_limiter.REDIS {
				limiter = rate_limiter.NewRedisRateLimiter(*cfg)
			}

			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				http.Error(w, "Internal error", http.StatusInternalServerError)
				return
			}

			token := r.Header.Get("API_KEY")

			if token != "" {
				if !limiter.AllowToken(token, cfg.RateLimitToken, time.Duration(cfg.BlockTimeRateLimitToken)*time.Second) {
					log.Println("Too Many Requests (Token Limit)")
					http.Error(w, "You have reached the maximum number of requests or actions allowed within a certain time frame.", http.StatusTooManyRequests)
					return
				}
			} else {
				if !limiter.AllowIP(ip, cfg.RateLimit, time.Duration(cfg.BlockTimeRateLimit)*time.Second) {
					log.Println("Too Many Requests (IP Limit)")
					http.Error(w, "You have reached the maximum number of requests or actions allowed within a certain time frame.", http.StatusTooManyRequests)
					return
				}
			}

			defer func() {
				if err := recover(); err != nil {
					log.Println("recovered from panic:", err)
					debug.PrintStack()
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
