package services

import (
	"encoding/json"
	"fmt"
	"github.com/leonardopinho/GoLang/4.Weather-API/cmd/domain"
	"github.com/leonardopinho/GoLang/4.Weather-API/config"
	"github.com/leonardopinho/GoLang/4.Weather-API/internal/entity"
	"net/http"
	"strings"
)

type OpenWeatherAPIResponse struct {
	Weather        *entity.Weather
	WeatherDetails *entity.WeatherDetails
}

type OpenWeatherService struct {
	Config *config.Config
}

func NewOpenWeatherService(cfg *config.Config) *OpenWeatherService {
	return &OpenWeatherService{
		Config: cfg,
	}
}

func (w OpenWeatherService) GetTemperature(location string) (*OpenWeatherAPIResponse, *domain.APIError) {
	location = strings.Replace(location, " ", "%20", -1)

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&lang=pt", w.Config.WeatherApiKey, location)
	resp, err := http.Get(url)
	if err != nil {
		return nil, domain.ErrWeatherServiceUnavailable
	}
	defer resp.Body.Close()

	var data OpenWeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data.Weather); err != nil {
		return nil, domain.ErrWeatherServiceInvalidResponse
	}
	data.WeatherDetails = entity.NewWeatherDetails(data.Weather.Current.TempC)

	return &data, nil
}
