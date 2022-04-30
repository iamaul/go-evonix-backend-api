package otel

import (
	"github.com/iamaul/go-evonix-backend-api/config"

	"go.opentelemetry.io/otel"
	jaegerExporter "go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

var Tracer = otel.Tracer("checkervisor-api")

func JaegerTelemetry(cfg *config.Config) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaegerExporter.New(
		jaegerExporter.WithCollectorEndpoint(jaegerExporter.WithEndpoint(cfg.Jaeger.Host)))
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.Jaeger.ServiceName),
			semconv.ServiceVersionKey.String(cfg.Server.AppVersion),
		)),
	)

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	return tp, nil
}
