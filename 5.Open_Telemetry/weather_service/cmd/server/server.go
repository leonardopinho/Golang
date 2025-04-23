package server

import (
	"context"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/weather_service/config"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/weather_service/infra/web"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Init() {
	cfg, err := config.LoadConfig("../")
	if err != nil {
		panic(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	shutdownTracer, err := config.InitProvider(ctx, cfg)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := shutdownTracer(context.Background()); err != nil {
			panic(err)
		}
	}()

	cfg.Tracer = otel.Tracer("weather_microservice_tracer")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Handle("/metrics", promhttp.Handler())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	r.Route("/get_address", func(r chi.Router) {
		r.Get("/{zip_code}", handlers.WeatherHandler(cfg))
	})

	var srv = &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	serverErr := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	select {
	case err := <-serverErr:
		log.Fatal(err)
		return

	case <-ctx.Done():
		stop()

		tracer := otel.Tracer("shutdown_tracer")
		_, span := tracer.Start(context.Background(), "GracefulShutdown")
		defer span.End()

		ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctxShutdown); err != nil {
			log.Fatalf("Erro no shutdown: %v", err)
		}

		log.Println("Server shutdown successfully.")
	}

	//
	//cfg, err := config.LoadConfig(".")
	//if err != nil {
	//	panic(err)
	//}
	//
	//ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	//defer stop()
	//
	//shutdownTracer, err := config.InitProvider(ctx, config.OTELConfig{
	//	ServiceName:  os.Getenv("OTEL_SERVICE_NAME"),
	//	CollectorURL: os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"),
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer func() {
	//	if err := shutdownTracer(context.Background()); err != nil {
	//		panic(err)
	//	}
	//}()
	//
	//r := chi.NewRouter()
	//r.Use(middleware.RequestID)
	//r.Use(middleware.RealIP)
	//r.Use(middleware.Recoverer)
	//r.Use(middleware.Logger)
	//r.Use(middleware.Timeout(60 * time.Second))
	//
	//r.Handle("/metrics", promhttp.Handler())
	//
	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	//r.Get("/get_weather", handlers.WeatherHandler(cfg))
	//
	//var srv = &http.Server{
	//	Addr:    ":8080",
	//	Handler: r,
	//}
	//
	//serverErr := make(chan error, 1)
	//go func() {
	//	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	//		serverErr <- err
	//	}
	//}()
	//
	//select {
	//case err := <-serverErr:
	//	log.Fatal(err)
	//	return
	//
	//case <-ctx.Done():
	//	stop()
	//
	//	tracer := otel.Tracer("shutdown-tracer")
	//	_, span := tracer.Start(context.Background(), "GracefulShutdown")
	//	defer span.End()
	//
	//	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//	defer cancel()
	//
	//	if err := srv.Shutdown(ctxShutdown); err != nil {
	//		log.Fatalf("Erro no shutdown: %v", err)
	//	}
	//
	//	log.Println("Server shutdown successfully.")
	//}
}
