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
	instance *rate_limiter.MemoryRateLimiter
)

func RateLimiterMiddleware(cfg *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			instance = rate_limiter.NewInMemoryRateLimiter(*cfg)

			token := r.Header.Get("API_KEY")
			if !instance.AllowToken(token, cfg.RateLimitToken, time.Duration(cfg.BlockTimeRateLimitToken)*time.Second) {
				http.Error(w, "Too Many Requests (Token Limit)", http.StatusTooManyRequests)
				return
			}

			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				http.Error(w, "Internal error", http.StatusInternalServerError)
				return
			}

			if !instance.AllowIP(ip, cfg.RateLimit, time.Duration(cfg.BlockTimeRateLimit)*time.Second) {
				http.Error(w, "Too Many Requests (IP Limit)", http.StatusTooManyRequests)
				return
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
