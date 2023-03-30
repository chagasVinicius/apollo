// Package opentel provides support for create opentelemetry tracer provider
// with jaeger protocol support.
package opentel

import (
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

// StartTracing configure open telemetry with jaeger support.
func StartTracing(serviceName string, reporterURI string, probability float64) (*trace.TracerProvider, error) {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(reporterURI)))
	if err != nil {
		return nil, fmt.Errorf("creating new exporter: %w", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.TraceIDRatioBased(probability)),
		// Always be sure to batch in production.
		trace.WithBatcher(exporter,
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
			trace.WithBatchTimeout(trace.DefaultScheduleDelay*time.Millisecond),
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
		),
		// Record information about this application in a Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
			attribute.String("exporter", "jaeger"),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}
