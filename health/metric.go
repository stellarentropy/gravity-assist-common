// Code generated by go generate; DO NOT EDIT.

package health

import (
	"context"
	"strings"

	t "github.com/stellarentropy/gravity-assist-common/metrics/tracer"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

func tracerAddInt64(ctx context.Context, name string, value int64, opts ...metric.AddOption) error {
	return t.AddInt64(ctx, componentName, name, value, opts...)
}

func tracerNewSpan(ctx context.Context, name string, attributes ...attribute.KeyValue) trace.Span {
	return t.NewSpan(ctx, strings.Join([]string{componentName, name}, "."), attributes...)
}

func getComponentFunctionName(name string) string {
	return strings.Join([]string{componentName, name}, ".")
}