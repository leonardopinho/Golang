package handlers

import (
	"encoding/json"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/weather_service/config"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/weather_service/internal/services"
	"net/http"
)

func WeatherHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//cep := r.URL.Query().Get("cep")
		//
		//cepService := services.NewViaCepService(cfg)
		//location, err := cepService.GetLocation(cep)
		//if err != nil {
		//	http.Error(w, err.Message, err.Status)
		//	return
		//}

		weatherService := services.NewOpenWeatherService(cfg)
		temp, err := weatherService.GetTemperature(location.Localidade)
		if err != nil {
			http.Error(w, err.Message, err.Status)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(&temp.WeatherDetails); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
