// Package validator provides gRPC interceptors for validating messages.
package validator

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/saltfishpr/go-learning/pkg/errors"
)

// The validateAller interface at protoc-gen-validate main branch.
// See https://github.com/envoyproxy/protoc-gen-validate/pull/468.
type validateAller interface {
	ValidateAll() error
}

// The validate interface starting with protoc-gen-validate v0.6.0.
// See https://github.com/envoyproxy/protoc-gen-validate/pull/455.
type validator interface {
	Validate(all bool) error
}

// The validate interface prior to protoc-gen-validate v0.6.0.
type validatorLegacy interface {
	Validate() error
}

func validate(req interface{}, all bool) error {
	if all {
		switch v := req.(type) {
		case validateAller:
			if err := v.ValidateAll(); err != nil {
				return err
			}
		case validator:
			if err := v.Validate(true); err != nil {
				return err
			}
		case validatorLegacy:
			// Fallback to legacy validator
			if err := v.Validate(); err != nil {
				return err
			}
		}
		return nil
	}
	switch v := req.(type) {
	case validatorLegacy:
		if err := v.Validate(); err != nil {
			return err
		}
	case validator:
		if err := v.Validate(false); err != nil {
			return err
		}
	}
	return nil
}

// UnaryServerInterceptor returns a new unary server interceptor that validates incoming messages.
//
// Invalid messages will be rejected with `InvalidArgument` before reaching any userspace handlers.
// If `all` is false, the interceptor returns first validation error. Otherwise the interceptor
// returns ALL validation error as a wrapped multi-error.
// Note that generated codes prior to protoc-gen-validate v0.6.0 do not provide an all-validation
// interface. In this case the interceptor fallbacks to legacy validation and `all` is ignored.
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := evaluateOptions(opts...)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := validate(req, o.all); err != nil {
			return nil, wrapError(err)
		}
		return handler(ctx, req)
	}
}

// UnaryClientInterceptor returns a new unary client interceptor that validates outgoing messages.
//
// Invalid messages will be rejected with `InvalidArgument` before sending the request to server.
// If `all` is false, the interceptor returns first validation error. Otherwise the interceptor
// returns ALL validation error as a wrapped multi-error.
// Note that generated codes prior to protoc-gen-validate v0.6.0 do not provide an all-validation
// interface. In this case the interceptor fallbacks to legacy validation and `all` is ignored.
func UnaryClientInterceptor(opts ...Option) grpc.UnaryClientInterceptor {
	o := evaluateOptions(opts...)
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if err := validate(req, o.all); err != nil {
			return wrapError(err)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that validates incoming messages.
//
// If `all` is false, the interceptor returns first validation error. Otherwise the interceptor
// returns ALL validation error as a wrapped multi-error.
// Note that generated codes prior to protoc-gen-validate v0.6.0 do not provide an all-validation
// interface. In this case the interceptor fallbacks to legacy validation and `all` is ignored.
// The stage at which invalid messages will be rejected with `InvalidArgument` varies based on the
// type of the RPC. For `ServerStream` (1:m) requests, it will happen before reaching any userspace
// handlers. For `ClientStream` (n:1) or `BidiStream` (n:m) RPCs, the messages will be rejected on
// calls to `stream.Recv()`.
func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	o := evaluateOptions(opts...)
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		wrapper := &recvWrapper{
			all:          o.all,
			ServerStream: stream,
		}
		return handler(srv, wrapper)
	}
}

type recvWrapper struct {
	all bool
	grpc.ServerStream
}

func (s *recvWrapper) RecvMsg(m interface{}) error {
	if err := s.ServerStream.RecvMsg(m); err != nil {
		return err
	}
	if err := validate(m, s.all); err != nil {
		return wrapError(err)
	}
	return nil
}

type validationError interface {
	Field() string
	Reason() string
	Cause() error
}

func extractCause(err error) error {
	if e, ok := err.(validationError); ok {
		if e.Cause() != nil {
			return extractCause(e.Cause())
		}
	}
	return err
}

func wrapError(err error) error {
	if e, ok := extractCause(err).(validationError); ok {
		return errors.New(codes.InvalidArgument, codes.InvalidArgument.String()).
			WithCause(err).
			WithMetadataPair(e.Field(), e.Reason())
	}
	return errors.New(codes.InvalidArgument, codes.InvalidArgument.String()).WithCause(err)
}
