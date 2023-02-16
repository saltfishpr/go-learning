package logging

import (
	"context"
	"fmt"
	"time"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	pkgerrors "github.com/pkg/errors"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"

	"github.com/saltfishpr/go-learning/pkg/errors"
)

var DefaultServerOptions = []grpc_zap.Option{
	grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
		return zap.String("grpc.latency", duration.String())
	}),
	grpc_zap.WithMessageProducer(func(ctx context.Context, msg string, level zapcore.Level, code codes.Code, err error, duration zapcore.Field) {
		fields := []zapcore.Field{
			zap.Uint32("grpc.code", uint32(code)),
			duration,
			zap.Error(err),
		}

		stack := extractStack(err)
		if stack != "" {
			fields = append(fields, zap.String("stack", stack))
		}

		ctxzap.Extract(ctx).Check(level, msg).Write(fields...)
	}),
}

func extractStack(err error) string {
	if err == nil {
		return ""
	}

	se, ok := lo.ErrorsAs[*errors.Error](err)
	if ok {
		st := errors.StackTrace(se.Unwrap())
		stack := getStack(st, 10, 2)
		return fmt.Sprintf("%s%+v", se.Unwrap().Error(), stack)
	}
	return ""
}

func getStack(st pkgerrors.StackTrace, depth int, skip int) pkgerrors.StackTrace {
	if len(st) < skip {
		return nil
	}

	if len(st) < depth {
		return st[skip:]
	}

	return st[skip : skip+depth]
}
