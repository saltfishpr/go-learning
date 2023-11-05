package internal

import (
	"context"
	"log/slog"
)

var ctxKeyLogger int

var defaultLogger = slog.Default()

func LoggerToContext(ctx context.Context, logger *slog.Logger) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, &ctxKeyLogger, logger)
}

func LoggerFromContext(ctx context.Context) *slog.Logger {
	if ctx == nil {
		return defaultLogger
	}
	logger, ok := ctx.Value(&ctxKeyLogger).(*slog.Logger)
	if !ok {
		return defaultLogger
	}
	return logger
}
