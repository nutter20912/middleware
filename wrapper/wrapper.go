package wrapper

import (
	"context"
	"fmt"
	"middleware/logging"
	"middleware/otel/tracer"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

type GinContextKey struct{}

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GinContextKey{}, c)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinContextKey{})
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}

	return gc, nil
}

func TracerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := otel.GetTextMapPropagator().Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

		opts := []trace.SpanStartOption{trace.WithSpanKind(trace.SpanKindServer)}
		ctx, span := tracer.StartSpanFromContext(ctx, c.Request.RequestURI, opts...)
		defer span.End()

		c.Request = c.Request.WithContext(ctx)

		logging.RequestLog(ctx, c.Request)

		c.Next()

		span.SetAttributes(
			semconv.HTTPStatusCode(c.Writer.Status()))

		if len(c.Errors) > 0 {
			span.SetStatus(codes.Error, "gin.errors")
			for _, err := range c.Errors {
				span.RecordError(err)
			}
		}
	}
}
