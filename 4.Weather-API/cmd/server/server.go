package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/leonardopinho/GoLang/4.Weather-API/config"
	"github.com/leonardopinho/GoLang/4.Weather-API/infra/web"
	"net/http"
)

func Init() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})

	r.Get("/get_weather", handlers.WeatherHandler(cfg))

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
