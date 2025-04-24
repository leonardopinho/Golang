package services

import (
	"context"
	"encoding/json"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/zip_code_service/cmd/domain"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/zip_code_service/config"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/zip_code_service/internal/entity"
	"net/http"
	"regexp"
)

type ZipCodeService struct {
	Config *config.Config
}

func NewZipCodeService(cfg *config.Config) *ZipCodeService {
	return &ZipCodeService{
		Config: cfg,
	}
}

func (v *ZipCodeService) IsValid(cep string) bool {
	match, _ := regexp.MatchString(`^\d{8}$`, cep)
	if cep == "" || len(cep) < 8 || !match {
		return false
	}
	return true
}

func (v *ZipCodeService) GetLocation(ctx context.Context, zip_code string) (*entity.WeatherDetail, *domain.APIError) {
	ctx, span := v.Config.Tracer.Start(ctx, "zip_code_service: GetLocation")
	defer span.End()

	if !v.IsValid(zip_code) {
		return nil, domain.ErrInvalidZipcode
	}

	url := v.Config.CepServiceURL + "/ws/" + zip_code + "/json/"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, domain.ErrNotFoundZipcode
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, domain.ErrCEPInvalidResponse
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, domain.ErrNotFoundZipcode
	}

	var address entity.Address
	if err = json.NewDecoder(resp.Body).Decode(&address); err != nil {
		return nil, domain.ErrCEPInvalidResponse
	}

	weather_service := NewWeatherService(v.Config)
	weather, api_err := weather_service.GetWeatherByLocation(address, ctx)
	if api_err != nil {
		return nil, api_err
	}

	weather_detail := entity.NewWeatherDetail(weather)
	return weather_detail, nil
}
