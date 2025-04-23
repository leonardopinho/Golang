package config

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/otel/trace"
	"log"
	"os"
	"path/filepath"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracer "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	CepServiceURL string
	OTELConfig    OTELConfig
	Tracer        trace.Tracer
}

type OTELConfig struct {
	ServiceName  string
	CollectorURL string
}

func LoadConfig(path string) (*Config, error) {
	envPath := filepath.Join(path, ".", ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Could not load .env: %v", err)
	}

	cepServiceURL := os.Getenv("VIA_CEP_SERVICE")

	config := &Config{
		CepServiceURL: cepServiceURL,
		OTELConfig: OTELConfig{
			ServiceName:  os.Getenv("OTEL_SERVICE_NAME"),
			CollectorURL: os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"),
		},
	}

	return config, nil
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
