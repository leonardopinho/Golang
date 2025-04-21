package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	WeatherApiKey     string
	WeatherServiceURL string
	CepServiceURL     string
}

func LoadConfig(path string) (*Config, error) {
	envPath := filepath.Join(path, ".", ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Could not load .env: %v", err)
	}

	weatherApiKey := os.Getenv("OPEN_WEATHERMAP_SERVICE_API_KEY")
	weatherServiceURL := os.Getenv("OPEN_WEATHERMAP_SERVICE")
	cepServiceURL := os.Getenv("VIA_CEP_SERVICE")

	config := &Config{
		WeatherApiKey:     weatherApiKey,
		CepServiceURL:     cepServiceURL,
		WeatherServiceURL: weatherServiceURL,
	}

	return config, nil
}
