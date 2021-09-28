// @description: 在上下文中存储与获取 logger
// @file: context.go
// @date: 2021/11/16

package logger

import (
	"context"

	"go.uber.org/zap"
)

type key int

const (
	loggerKey key = iota
)

// NewContextWithLogger returns a new context with log added.
func NewContextWithLogger(ctx context.Context, log *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerKey, log)
}

// FromContext returns the *zap.SugaredLogger associated with ctx or nil if no logger has been assigned.
func FromContext(ctx context.Context) *zap.SugaredLogger {
	l, _ := ctx.Value(loggerKey).(*zap.SugaredLogger)
	return l
}
