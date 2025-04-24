package domain

import "net/http"

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

func (a APIError) Error() string {
	return a.Message
}

var (
	ErrInvalidZipcode = &APIError{
		Code:    "INVALID_ZIPCODE",
		Message: "invalid zipcode",
		Status:  http.StatusUnprocessableEntity,
	}

	ErrNotFoundZipcode = &APIError{
		Code:    "NOT_FOUND_ZIPCODE",
		Message: "can not find zipcode",
		Status:  http.StatusNotFound,
	}

	ErrCEPInvalidResponse = &APIError{
		Code:    "CEP_INVALID_RESPONSE",
		Message: "failed to process address data",
		Status:  http.StatusInternalServerError,
	}

	ErrInvalidResponse = &APIError{
		Code:    "INVALID_RESPONSE",
		Message: "failed to process data",
		Status:  http.StatusInternalServerError,
	}

	ErrWeatherServiceInvalidResponse = &APIError{
		Code:    "WEATHER_INVALID_RESPONSE",
		Message: "failed to process weather data",
		Status:  http.StatusInternalServerError,
	}

	ErrWeatherServiceUnavailable = &APIError{
		Code:    "WEATHER_SERVICE_UNAVAILABLE",
		Message: "failed to send data",
		Status:  http.StatusInternalServerError,
	}
)
