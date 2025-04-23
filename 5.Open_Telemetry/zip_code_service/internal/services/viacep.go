package services

import (
	"context"
	"encoding/json"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/zip_code_service/cmd/domain"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/zip_code_service/config"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/zip_code_service/internal/entity"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
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

func (v *ViaCepService) GetLocation(cep string, ctx context.Context, tracer trace.Tracer) (*entity.Address, *domain.APIError) {
	ctx, span := tracer.Start(ctx, "get_zip_code")
	defer span.End()

	if !v.IsValid(cep) {
		return nil, domain.ErrInvalidZipcode
	}

	url := v.Config.CepServiceURL + "/ws/" + cep + "/json/"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, domain.ErrNotFoundZipcode
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, domain.ErrCEPInvalidResponse
	}
	defer resp.Body.Close()

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	if resp.StatusCode == http.StatusNotFound {
		return nil, domain.ErrNotFoundZipcode
	}

	var address entity.Address

	if err = json.NewDecoder(resp.Body).Decode(&address); err != nil {
		return nil, domain.ErrCEPInvalidResponse
	}

	return &address, nil
}
