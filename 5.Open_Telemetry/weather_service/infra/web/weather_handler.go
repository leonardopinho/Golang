package handlers

import (
	"encoding/json"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/weather_service/cmd/domain"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/weather_service/config"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/weather_service/internal/entity"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/weather_service/internal/services"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"net/http"
)

func WeatherHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var address entity.Address
		if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
			http.Error(w, domain.ErrWeatherServiceInvalidResponse.Message, domain.ErrWeatherServiceInvalidResponse.Status)
			return
		}

		if address.Cep == "" || address.Localidade == "" {
			http.Error(w, domain.ErrWeatherServiceInvalidResponse.Message, domain.ErrWeatherServiceInvalidResponse.Status)
			return
		}

		ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))

		weatherService := services.NewOpenWeatherService(cfg)
		temp, err := weatherService.GetTemperature(ctx, address.Localidade)
		if err != nil {
			http.Error(w, err.Message, err.Status)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(temp.Weather); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
