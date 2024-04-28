package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Salaton/tracing/pkg/presentation"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func startTracing() (*trace.TracerProvider, error) {
	serviceName := "product-app"
	headers := map[string]string{
		"content-type": "application/json",
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint("localhost:4318"),
			otlptracehttp.WithHeaders(headers),
			otlptracehttp.WithInsecure(),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("creating new exporter: %w", err)
	}

	tracerprovider := trace.NewTracerProvider(
		trace.WithBatcher(
			exporter,
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
			trace.WithBatchTimeout(trace.DefaultScheduleDelay*time.Millisecond),
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
		),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(serviceName),
				attribute.String("environment", "testing"),
			),
		),
	)

	otel.SetTracerProvider(tracerprovider)

	return tracerprovider, nil
}

func main() {
	traceProvider, err := startTracing()
	if err != nil {
		log.Fatalf("traceprovider: %v", err)
	}
	defer func() {
		if err := traceProvider.Shutdown(context.Background()); err != nil {
			log.Fatalf("traceprovider: %v", err)
		}
	}()

	_ = traceProvider.Tracer("my-app")

	presentation.StartServer(8080)

	// Block until we receive a sigint (CTRL+C) signal
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
}
