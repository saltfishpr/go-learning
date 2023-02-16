// Package server .
package server

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
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

	"github.com/saltfishpr/go-learning/internal/user/conf"
	"github.com/saltfishpr/go-learning/pkg/errors"
	_grpc_auth_jwt "github.com/saltfishpr/go-learning/pkg/interceptor/auth/jwt"
	_grpc_logging "github.com/saltfishpr/go-learning/pkg/interceptor/logging"
	_grpc_validator "github.com/saltfishpr/go-learning/pkg/interceptor/validator"
)

func NewGRPCServer(i *do.Injector) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(do.MustInvoke[*zap.Logger](i), loggingOptions...),
			grpc_zap.PayloadUnaryServerInterceptor(do.MustInvoke[*zap.Logger](i), unaryPayloadLoggingDecider()),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(recoverHandleFunc())),
			grpc_auth.UnaryServerInterceptor(_grpc_auth_jwt.AuthFunc(authFunc(do.MustInvoke[*conf.Config](i)))),
			_grpc_validator.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(do.MustInvoke[*zap.Logger](i), loggingOptions...),
			grpc_zap.PayloadStreamServerInterceptor(do.MustInvoke[*zap.Logger](i), streamPayloadLoggingDecider()),
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(recoverHandleFunc())),
			grpc_auth.StreamServerInterceptor(_grpc_auth_jwt.AuthFunc(authFunc(do.MustInvoke[*conf.Config](i)))),
			_grpc_validator.StreamServerInterceptor(),
		)),
	)

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
		return errors.New(errors.NewStatus(codes.Unknown, codes.Unknown.String())).
			WithCause(fmt.Errorf("%v", p))
	}
}

func authFunc(config *conf.Config) func(tokenStr string) (*jwt.MapClaims, error) {
	return func(tokenStr string) (*jwt.MapClaims, error) {
		return nil, nil
	}
}
