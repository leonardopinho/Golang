package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/leonardopinho/GoLang/2.Rate_Limiter/config"
	handlers "github.com/leonardopinho/GoLang/2.Rate_Limiter/internal/infra/webserver/handlers"
	customMiddleware "github.com/leonardopinho/GoLang/2.Rate_Limiter/internal/middleware"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(customMiddleware.RateLimiterMiddleware(cfg))
	r.Use(middleware.Recoverer)

	r.Get("/", handlers.IndexHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
