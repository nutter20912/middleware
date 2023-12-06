package logging

import (
	"context"
	"encoding/json"
	"net/http"

	"go-micro.dev/v4/client"
	"go-micro.dev/v4/metadata"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

func RequestLog(ctx context.Context, req *http.Request) {
	span := trace.SpanFromContext(ctx)

	attrs := []attribute.KeyValue{
		semconv.HTTPURL(req.URL.String()),
		semconv.HTTPMethod(req.Method),
		semconv.ClientAddress(req.RemoteAddr)}

	span.SetAttributes(attrs...)

	payload, _ := json.Marshal(map[string]interface{}{
		"header": req.Header,
	})

	span.AddEvent("request", trace.WithAttributes(
		attribute.String("payload", string(payload))))
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
