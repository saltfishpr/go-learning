// Package selector .
package selector

import (
	"context"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

type MatchFunc func(ctx context.Context, name string) bool

func UnaryServer(opts ...Option) func(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	o := evaluateOptions(opts...)
	return func(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
		chaind := grpc_middleware.ChainUnaryServer(interceptors...)
		return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
			if o.match(ctx, info.FullMethod) {
				return chaind(ctx, req, info, handler)
			}
			return handler(ctx, req)
		}
	}
}

func StreamServer(opts ...Option) func(interceptors ...grpc.StreamServerInterceptor) grpc.StreamServerInterceptor {
	o := evaluateOptions(opts...)
	return func(interceptors ...grpc.StreamServerInterceptor) grpc.StreamServerInterceptor {
		chaind := grpc_middleware.ChainStreamServer(interceptors...)
		return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			if o.match(ss.Context(), info.FullMethod) {
				return chaind(srv, ss, info, handler)
			}
			return handler(srv, ss)
		}
	}
}
