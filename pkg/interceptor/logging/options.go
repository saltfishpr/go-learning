package logging

import (
	"context"
	"time"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
)

var DefaultServerOptions = []grpc_zap.Option{
	grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
		return zap.String("grpc.latency", duration.String())
	}),
	grpc_zap.WithMessageProducer(
		func(ctx context.Context, msg string, level zapcore.Level, code codes.Code, err error, duration zapcore.Field) {
			fields := []zapcore.Field{
				zap.Uint32("grpc.code", uint32(code)),
				duration,
				zap.Error(err),
			}

			ctxzap.Extract(ctx).Check(level, msg).Write(fields...)
		},
	),
}
