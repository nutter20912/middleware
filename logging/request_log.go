package logging

import (
	"context"
	"encoding/json"

	"go-micro.dev/v4/client"
	"go-micro.dev/v4/metadata"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func RequestLog() {

}

func ClientRequestLog(ctx context.Context, req client.Request, rsp interface{}) {
	span := trace.SpanFromContext(ctx)

	attrs := []attribute.KeyValue{}

	span.SetAttributes(attrs...)

	md, _ := metadata.FromContext(ctx)
	payload, _ := json.Marshal(map[string]interface{}{
		"metadata":    md,
		"service":     req.Service(),
		"method":      req.Method(),
		"contentType": req.ContentType(),
		"body":        req.Body(),
	})

	span.AddEvent("request", trace.WithAttributes(
		attribute.String("payload", string(payload))))

	rspBytes, _ := json.Marshal(rsp)
	span.AddEvent("response",
		trace.WithAttributes(attribute.String("response.message", string(rspBytes))))

}
