package tests

import (
	"context"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/zip_code_service/config"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/zip_code_service/internal/services"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel"
	"testing"
)

func TestGetLocation(t *testing.T) {
	cfg := &config.Config{
		CepServiceURL:     "https://viacep.com.br",
		Tracer:            otel.Tracer("test_tracer"),
		WeatherServiceURL: "http://127.0.0.1:8081/get_weather",
		OTELConfig: config.OTELConfig{
			ServiceName:  "",
			CollectorURL: "",
		},
	}

	ctx := context.Background()
	cep := "13330250"
	cepService := services.NewZipCodeService(cfg)
	location, _ := cepService.GetLocation(ctx, cep)

	assert.Equal(t, location.City, "Indaiatuba", true)
}
