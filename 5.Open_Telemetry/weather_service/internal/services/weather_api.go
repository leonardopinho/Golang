package services

import (
	"encoding/json"
	"fmt"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/weather_service/cmd/domain"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/weather_service/config"
	"github.com/leonardopinho/GoLang/5.Open_Telemetry/weather_service/internal/entity"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracer "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strings"
)

type OpenWeatherAPIResponse struct {
	Weather        *entity.Weather
	WeatherDetails *entity.WeatherDetails
}

type OpenWeatherService struct {
	Config *config.Config
}

func NewOpenWeatherService(cfg *config.Config) *OpenWeatherService {
	return &OpenWeatherService{
		Config: cfg,
	}
}

func InitProvider(ctx context.Context, cfg *Config) (func(context.Context) error, error) {
	if cfg.OTELConfig.ServiceName == "" {
		cfg.OTELConfig.ServiceName = "default-service"
	}
	if cfg.OTELConfig.CollectorURL == "" {
		cfg.OTELConfig.CollectorURL = "localhost:4317"
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(cfg.OTELConfig.ServiceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	conn, err := grpc.NewClient(cfg.OTELConfig.CollectorURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create gRPC connection. service=%q collector=%q: %w",
			cfg.OTELConfig.ServiceName, cfg.OTELConfig.CollectorURL, err,
		)
	}

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	bsp := tracer.NewBatchSpanProcessor(traceExporter)

	tracerProvider := tracer.NewTracerProvider(
		tracer.WithSampler(tracer.AlwaysSample()),
		tracer.WithResource(res),
		tracer.WithSpanProcessor(bsp),
	)

	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(tracerProvider)

	return tracerProvider.Shutdown, nil
}
