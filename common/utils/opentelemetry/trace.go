package opentelemetry

import (
	"encoding/json"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func Log(span trace.Span, input interface{}, output interface{}, err error) {
	inputByte, _ := json.Marshal(input)
	outputByte, _ := json.Marshal(output)
	span.SetAttributes(attribute.String("input", string(inputByte)))
	span.SetAttributes(attribute.String("output", string(outputByte)))

	if err != nil {
		span.SetAttributes(attribute.String("error", err.Error()))
	}

	span.End()
}
