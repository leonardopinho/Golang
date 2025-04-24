package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/weather_service/cmd/domain"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/weather_service/config"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/weather_service/internal/entity"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"net/http"
	"strings"
)

type WeatherAPIResponse struct {
	Weather        *entity.Weather
	WeatherDetails *entity.WeatherDetails
}

type WeatherService struct {
	Config *config.Config
}

func NewOpenWeatherService(cfg *config.Config) *WeatherService {
	return &WeatherService{
		Config: cfg,
	}
}

func (w WeatherService) GetTemperature(ctx context.Context, location string) (*WeatherAPIResponse, *domain.APIError) {
	ctx, span := w.Config.Tracer.Start(ctx, "weather_service: GetTemperature")
	defer span.End()

	location = strings.Replace(location, " ", "%20", -1)
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&lang=pt", w.Config.WeatherApiKey, location)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, domain.ErrWeatherServiceUnavailable
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, domain.ErrWeatherServiceUnavailable
	}
	defer resp.Body.Close()

	var data WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data.Weather); err != nil {
		return nil, domain.ErrWeatherServiceInvalidResponse
	}

	data.WeatherDetails = entity.NewWeatherDetails(data.Weather)
	return &data, nil
}
