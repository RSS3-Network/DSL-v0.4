package opentelemetry

import (
	"os"

	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

const (
	DefaultPath = "traces.log"
)

func DialWithURL(endpointURL string) (exporter trace.SpanExporter, err error) {
	if exporter, err = jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpointURL))); err != nil {
		return nil, err
	}

	return exporter, nil
}

func DialWithPath(path string) (exporter trace.SpanExporter, err error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	exporter, err = stdouttrace.New(
		stdouttrace.WithWriter(file),
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithoutTimestamps(),
	)

	return exporter, nil
}
