package log

import (
	"context"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

// Panicf logs a message at level Panic on the standard logger.
func Panicf(ctx context.Context, format string, args ...interface{}) {
	loggerFromContext(ctx).Panicf(addTraceID(ctx, format), args...)
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	loggerFromContext(ctx).Fatalf(addTraceID(ctx, format), args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(ctx context.Context, format string, args ...interface{}) {
	loggerFromContext(ctx).Errorf(addTraceID(ctx, format), args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(ctx context.Context, format string, args ...interface{}) {
	loggerFromContext(ctx).Warnf(addTraceID(ctx, format), args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(ctx context.Context, format string, args ...interface{}) {
	loggerFromContext(ctx).Infof(addTraceID(ctx, format), args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(ctx context.Context, format string, args ...interface{}) {
	loggerFromContext(ctx).Debugf(addTraceID(ctx, format), args...)
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(ctx context.Context, format string, args ...interface{}) {
	loggerFromContext(ctx).Tracef(addTraceID(ctx, format), args...)
}

func addTraceID(ctx context.Context, format string) string {
	traceID := newrelic.FromContext(ctx)
	return format + " [trace_id:" + traceID.GetTraceMetadata().TraceID + "]"
}

type logger struct{}

// WithLogger returns a copy of parent in which the standard logger is replaced.
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
