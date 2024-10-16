package pinduoduo

import (
	"context"
	"github.com/mazesoul87/go-library/gorequest"
	"go.opentelemetry.io/otel/trace"
)

// TraceStartSpan 开始OpenTelemetry链路追踪状态
func TraceStartSpan(ctx context.Context, spanName string) (context.Context, trace.Span) {
	return gorequest.TraceNewSpan(ctx, "github.com/mazesoul87/go-library/service/pinduoduo", "pinduoduo.", spanName, Version, trace.SpanKindClient)
}
