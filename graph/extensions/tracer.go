package extensions

import (
	"context"
	"middleware/logging"
	"middleware/otel/tracer"

	"github.com/99designs/gqlgen/graphql"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type Tracer struct{}

func NewTracer() Tracer {
	return Tracer{}
}

func (t Tracer) ExtensionName() string {
	return "OtelTracer"
}

func (t Tracer) Validate(_ graphql.ExecutableSchema) error {
	return nil
}

func (t Tracer) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	if !graphql.HasOperationContext(ctx) {
		return next(ctx)
	}

	oc := graphql.GetOperationContext(ctx)

	opts := []trace.SpanStartOption{trace.WithSpanKind(trace.SpanKindInternal)}

	ctx, span := tracer.StartSpanFromContext(ctx, oc.OperationName, opts...)

	logging.OperationLog(ctx, oc)

	rspHander := next(ctx)

	return func(ctx context.Context) *graphql.Response {
		defer span.End()

		return rspHander(ctx)
	}
}

func (t Tracer) InterceptField(ctx context.Context, next graphql.Resolver) (interface{}, error) {
	fc := graphql.GetFieldContext(ctx)

	if fc == nil || !fc.IsResolver {
		return next(ctx)
	}

	opts := []trace.SpanStartOption{trace.WithSpanKind(trace.SpanKindInternal)}

	operationName := fc.Field.ObjectDefinition.Name + "/" + fc.Field.Name
	ctx, span := tracer.StartSpanFromContext(ctx, operationName, opts...)
	defer span.End()

	logging.FieldLog(ctx, fc)

	rsp, err := next(ctx)
	if err != nil {
		span.SetStatus(codes.Error, "field err")
		span.RecordError(err)
	}

	return rsp, err
}

func (t Tracer) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	span := trace.SpanFromContext(ctx)

	rsp := next(ctx)
	if rsp != nil && len(rsp.Errors) != 0 {
		span.SetStatus(codes.Error, "response err")
		for _, err := range rsp.Errors {
			span.RecordError(err)
		}
	}

	return rsp
}
