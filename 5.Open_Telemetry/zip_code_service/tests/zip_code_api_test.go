package tests

import (
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/zip_code_service/config"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/zip_code_service/internal/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLocation(t *testing.T) {
	cfg := &config.Config{
		CepServiceURL: "https://viacep.com.br",
	}

	cep := "13330250"
	cepService := services.NewViaCepService(cfg)
	location, _ := cepService.GetLocation(cep, nil, nil)

	assert.Equal(t, location.Uf, "SP", true)
}
