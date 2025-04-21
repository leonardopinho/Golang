package services

import (
	"encoding/json"
	"github.com/leonardopinho/GoLang/4.Weather-API/cmd/domain"
	"github.com/leonardopinho/GoLang/4.Weather-API/config"
	"github.com/leonardopinho/GoLang/4.Weather-API/internal/entity"
	"io"
	"net/http"
	"regexp"
)

type ViaCepService struct {
	Config *config.Config
}

func NewViaCepService(cfg *config.Config) *ViaCepService {
	return &ViaCepService{
		Config: cfg,
	}
}

func (v *ViaCepService) IsValid(cep string) bool {
	match, _ := regexp.MatchString(`^\d{8}$`, cep)
	if cep == "" || len(cep) < 8 || !match {
		return false
	}
	return true
}

func (v *ViaCepService) GetLocation(cep string) (*entity.Address, *domain.APIError) {
	if !v.IsValid(cep) {
		return nil, domain.ErrInvalidZipcode
	}

	url := v.Config.CepServiceURL + "/ws/" + cep + "/json/"
	req, err := http.Get(url)
	if err != nil {
		return nil, domain.ErrNotFoundZipcode
	}
	defer req.Body.Close()

	if req.StatusCode == http.StatusNotFound {
		return nil, domain.ErrNotFoundZipcode
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, domain.ErrCEPInvalidResponse
	}

	var address entity.Address
	if err = json.Unmarshal(res, &address); err != nil {
		return nil, domain.ErrCEPInvalidResponse
	}

	return &address, nil
}
