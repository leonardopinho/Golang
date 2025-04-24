package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/zip_code_service/config"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/zip_code_service/internal/services"
	"net/http"
)

func ZipCodeHandler(cfg *config.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		zip_code := chi.URLParam(r, "zip_code")
		if zip_code == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ctx, span := cfg.Tracer.Start(ctx, cfg.OTELConfig.ServiceName)
		defer span.End()

		cepService := services.NewZipCodeService(cfg)
		location, err := cepService.GetLocation(ctx, zip_code)
		if err != nil {
			http.Error(w, err.Message, err.Status)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&location); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
