package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	RateLimit               int
	RateLimitToken          int
	BlockTimeRateLimit      int
	BlockTimeRateLimitToken int
	Strategy                int
	RedisAddr               string
	RedisPassword           string
	RedisDB                 int
}

func LoadConfig(path string) (*Config, error) {
	envPath := filepath.Join(path, ".", ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Could not load .env: %v", err)
	}

	rateLimit, err := strconv.Atoi(os.Getenv("RATE_LIMIT"))
	if err != nil {
		return nil, fmt.Errorf("error converting RATE_LIMIT: %w", err)
	}

	rateLimitToken, err := strconv.Atoi(os.Getenv("RATE_LIMIT_TOKEN"))
	if err != nil {
		return nil, fmt.Errorf("error converting RATE_LIMIT_TOKEN: %w", err)
	}

	blockTimeRateLimit, err := strconv.Atoi(os.Getenv("BLOCK_TIME_RATE_LIMIT"))
	if err != nil {
		return nil, fmt.Errorf("error converting BLOCK_TIME_RATE_LIMIT: %w", err)
	}

	blockTimeRateLimitToken, err := strconv.Atoi(os.Getenv("BLOCK_TIME_RATE_LIMIT_TOKEN"))
	if err != nil {
		return nil, fmt.Errorf("error converting BLOCK_TIME_RATE_LIMIT_TOKEN: %w", err)
	}

	strategy, err := strconv.Atoi(os.Getenv("STRATEGY"))
	if err != nil {
		return nil, fmt.Errorf("error converting STRATEGY: %w", err)
	}

	redisAddr := os.Getenv("REDIS_ADDR")

	redisPassword := os.Getenv("REDIS_PASSWORD")

	redisDb, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, fmt.Errorf("error converting REDIS_DB: %w", err)
	}

	config := &Config{
		RateLimit:               rateLimit,
		RateLimitToken:          rateLimitToken,
		BlockTimeRateLimit:      blockTimeRateLimit,
		BlockTimeRateLimitToken: blockTimeRateLimitToken,
		Strategy:                strategy,
		RedisAddr:               redisAddr,
		RedisPassword:           redisPassword,
		RedisDB:                 redisDb,
	}

	return config, nil
}
