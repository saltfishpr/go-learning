package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

// StructuredLogger implements the middleware.LogFormatter interface with zap
// logger.
type StructuredLogger struct {
	Logger *zap.Logger
}

var _ middleware.LogFormatter = (*StructuredLogger)(nil)

func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	var fields []zap.Field

	if requestID := GetRequestID(r.Context()); requestID != "" {
		fields = append(fields, zap.String("request_id", requestID))
	}
	fields = append(fields, (zap.String("method", r.Method)))
	fields = append(fields, (zap.String("path", r.RequestURI)))
	fields = append(fields, (zap.String("user_agent", r.UserAgent())))

	return &StructuredLogEntry{Logger: l.Logger.With(fields...)}
}

type StructuredLogEntry struct {
	Logger *zap.Logger
	fields []zap.Field // fields to be added to the logger
}

var _ middleware.LogEntry = (*StructuredLogEntry)(nil)

func (l *StructuredLogEntry) Write(
	status int, _ int, _ http.Header, elapsed time.Duration, _ interface{},
) {
	l.Logger = l.Logger.With(zap.Int("status", status))
	l.Logger = l.Logger.With(zap.String("elapsed", elapsed.String()))
	l.Logger = l.Logger.With(l.fields...)

	switch {
	case status >= 500:
		l.Logger.Error("request complete")
	default:
		l.Logger.Info("request complete")
	}
}

func (l *StructuredLogEntry) Panic(v interface{}, stack []byte) {
	l.fields = append(l.fields, zap.String("stack", string(stack)))
	l.fields = append(l.fields, zap.String("panic", fmt.Sprintf("%+v", v)))
}

func LogEntrySetFields(ctx context.Context, fields ...zap.Field) {
	if entry, ok := ctx.Value(middleware.LogEntryCtxKey).(*StructuredLogEntry); ok {
		entry.fields = append(entry.fields, fields...)
	}
}
