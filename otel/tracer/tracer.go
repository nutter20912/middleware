package tracer

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var (
	instrumentationName = "middleware/otel/tacer"
)

func StartSpanFromContext(ctx context.Context, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	ctx, span := otel.Tracer(instrumentationName).Start(ctx, name, opts...)
	return ctx, span
}
