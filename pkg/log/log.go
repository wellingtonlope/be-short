package log

import (
	"context"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

func Panicf(ctx context.Context, format string, args ...interface{}) {
	loggerFromContext(ctx).Panicf(addTraceID(ctx, format), args...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	loggerFromContext(ctx).Fatalf(addTraceID(ctx, format), args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	loggerFromContext(ctx).Errorf(addTraceID(ctx, format), args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	loggerFromContext(ctx).Warnf(addTraceID(ctx, format), args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	loggerFromContext(ctx).Infof(addTraceID(ctx, format), args...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	loggerFromContext(ctx).Debugf(addTraceID(ctx, format), args...)
}

func Tracef(ctx context.Context, format string, args ...interface{}) {
	loggerFromContext(ctx).Tracef(addTraceID(ctx, format), args...)
}

func addTraceID(ctx context.Context, format string) string {
	traceID := newrelic.FromContext(ctx)
	return format + " [trace_id:" + traceID.GetTraceMetadata().TraceID + "]"
}

type logger struct{}

func WithLogger(ctx context.Context, loggerLogrus *logrus.Logger) context.Context {
	return context.WithValue(ctx, logger{}, loggerLogrus)
}

func loggerFromContext(ctx context.Context) *logrus.Logger {
	loggerCtx, ok := ctx.Value(logger{}).(*logrus.Logger)
	if !ok {
		loggerCtx = logrus.StandardLogger()
	}
	return loggerCtx
}
