package tracer

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"log"
)

type TracerStruct struct {
	context context.Context
	logger  *log.Logger
	tracer  trace.Tracer
}

func NewTracer(context context.Context, tr trace.Tracer) *TracerStruct {
	return &TracerStruct{
		context: context,
		logger:  log.Default(),
		tracer:  tr,
	}
}

func (t *TracerStruct) TraceDecorator(fn GenericFuncType, functionName string) GenericFuncType {
	return func(args ...interface{}) (interface{}, error) {
		t.logger.Printf("Tracing function call, args: %s\n", functionName)
		_, tr := t.tracer.Start(t.context, functionName)
		defer tr.End()

		return fn(args...)
	}
}
