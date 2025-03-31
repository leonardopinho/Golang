package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/leonardopinho/GoLang/2.Rate_Limiter/config"
	handler "github.com/leonardopinho/GoLang/2.Rate_Limiter/internal/infra/webserver/handlers"
	custom_middleware "github.com/leonardopinho/GoLang/2.Rate_Limiter/internal/middleware"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(custom_middleware.RateLimiterMiddleware(cfg))
	r.Use(middleware.Recoverer)

	r.Get("/", handler.IndexHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
