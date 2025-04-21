package tests

import (
	"github.com/leonardopinho/GoLang/4.Weather-API/config"
	"github.com/leonardopinho/GoLang/4.Weather-API/internal/entity"
	"github.com/leonardopinho/GoLang/4.Weather-API/internal/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertTemperature(t *testing.T) {
	celcius := 17.9
	result := entity.ConvertTemperature(celcius)

	assert.Equal(t, result.Temp_F, 64.22, true)
	assert.Equal(t, result.Temp_K, 290.9, true)
}

func TestGetLocation(t *testing.T) {
	cfg := &config.Config{
		CepServiceURL: "https://viacep.com.br",
	}

	cep := "13330250"
	cepService := services.NewViaCepService(cfg)
	location, _ := cepService.GetLocation(cep)

	assert.Equal(t, location.Uf, "SP", true)
}
