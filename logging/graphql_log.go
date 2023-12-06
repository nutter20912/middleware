package logging

import (
	"context"
	"encoding/json"

	"github.com/99designs/gqlgen/graphql"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

func OperationLog(ctx context.Context, oc *graphql.OperationContext) {
	span := trace.SpanFromContext(ctx)

	attrs := []attribute.KeyValue{
		attribute.String("log_type", "operation_log"),
		semconv.GraphqlOperationName(oc.Operation.Name),
		semconv.GraphqlDocument(oc.RawQuery),
		attribute.String("operation_type", string(oc.Operation.Operation)),
	}

	span.SetAttributes(attrs...)

	payload, _ := json.Marshal(map[string]interface{}{
		"header": oc.Headers,
	})

	span.AddEvent("request", trace.WithAttributes(
		attribute.String("payload", string(payload))))
}

func FieldLog(ctx context.Context, fc *graphql.FieldContext) {
	span := trace.SpanFromContext(ctx)

	argsBytes, _ := json.Marshal(fc.Args)

	attrs := []attribute.KeyValue{
		attribute.String("log_type", "field_log"),
		attribute.String("object", fc.Object),
		attribute.String("name", fc.Field.Name),
		attribute.String("alias", fc.Field.Alias),
		attribute.String("path", fc.Path().String()),
		attribute.String("args", string(argsBytes)),
	}

	span.SetAttributes(attrs...)
}
