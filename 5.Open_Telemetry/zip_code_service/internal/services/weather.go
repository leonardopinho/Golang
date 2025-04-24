package services

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/zip_code_service/cmd/domain"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/zip_code_service/config"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/zip_code_service/internal/entity"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"net/http"
)

type WeatherService struct {
	Config     config.Config
	URLService string
}

func NewWeatherService(config *config.Config) *WeatherService {
	return &WeatherService{
		Config:     *config,
		URLService: config.WeatherServiceURL,
	}
}

func (w *WeatherService) GetWeatherByLocation(address entity.Address, ctx context.Context) (*entity.Weather, *domain.APIError) {
	payload, err := json.Marshal(address)
	if err != nil {
		return nil, domain.ErrInvalidResponse
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, w.Config.WeatherServiceURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, domain.ErrWeatherServiceUnavailable
	}

	req.Header.Set("Content-Type", "application/json")

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, domain.ErrWeatherServiceUnavailable
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, domain.ErrWeatherServiceInvalidResponse
	}

	var weather entity.Weather
	if err = json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, domain.ErrWeatherServiceInvalidResponse
	}

	return &weather, nil
}
