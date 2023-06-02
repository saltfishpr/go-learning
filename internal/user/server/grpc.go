// Package server .
package server

import (
	"context"
	stderrors "errors"
	"fmt"
	"time"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/samber/do"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"

	userv1 "github.com/saltfishpr/go-learning/gen/go/user/v1"
	"github.com/saltfishpr/go-learning/internal/user/conf"
	"github.com/saltfishpr/go-learning/internal/user/service"
	"github.com/saltfishpr/go-learning/pkg/errors"
	_grpc_logging "github.com/saltfishpr/go-learning/pkg/interceptor/logging"
	_grpc_validator "github.com/saltfishpr/go-learning/pkg/interceptor/validator"
)

func NewGRPC(i *do.Injector) *grpc.Server {
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: 5 * time.Minute,
		}),
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(do.MustInvoke[*zap.Logger](i), loggingOptions...),
			grpc_zap.PayloadUnaryServerInterceptor(do.MustInvoke[*zap.Logger](i), unaryPayloadLoggingDecider()),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(recoverHandleFunc())),
			grpc_auth.UnaryServerInterceptor(authFunc(do.MustInvoke[*conf.Config](i))),
			_grpc_validator.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(do.MustInvoke[*zap.Logger](i), loggingOptions...),
			grpc_zap.PayloadStreamServerInterceptor(do.MustInvoke[*zap.Logger](i), streamPayloadLoggingDecider()),
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(recoverHandleFunc())),
			grpc_auth.StreamServerInterceptor(authFunc(do.MustInvoke[*conf.Config](i))),
			_grpc_validator.StreamServerInterceptor(),
		),
	)

	userv1.RegisterUserServiceServer(s, do.MustInvoke[*service.UserService](i))

	return s
}

var loggingOptions = append([]grpc_zap.Option{
	grpc_zap.WithLevels(func(code codes.Code) zapcore.Level {
		switch code {
		case codes.OK:
			return zapcore.InfoLevel
		default:
			return zapcore.ErrorLevel
		}
	}),
}, _grpc_logging.DefaultServerOptions...)

func unaryPayloadLoggingDecider() grpc_logging.ServerPayloadLoggingDecider {
	return func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
		return true
	}
}

func streamPayloadLoggingDecider() grpc_logging.ServerPayloadLoggingDecider {
	return func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
		return true
	}
}

func recoverHandleFunc() grpc_recovery.RecoveryHandlerFuncContext {
	return func(ctx context.Context, p interface{}) (err error) {
		return errors.New(codes.Unknown, codes.Unknown.String()).
			WithCause(fmt.Errorf("%v", p))
	}
}

func authFunc(config *conf.Config) grpc_auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		return ctx, nil
	}
}

func getTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", stderrors.New("no metadata in context")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return "", stderrors.New("no authorization")
	}

	return values[0], nil
}
