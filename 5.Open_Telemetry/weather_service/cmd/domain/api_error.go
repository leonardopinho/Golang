package domain

import "net/http"

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

var (
	ErrWeatherServiceUnavailable = &APIError{
		Code:    "WEATHER_SERVICE_UNAVAILABLE",
		Message: "could not connect to weather service",
		Status:  http.StatusBadGateway,
	}

	ErrWeatherServiceInvalidResponse = &APIError{
		Code:    "WEATHER_INVALID_RESPONSE",
		Message: "failed to process weather data",
		Status:  http.StatusInternalServerError,
	}
)
