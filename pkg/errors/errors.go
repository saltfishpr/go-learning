// Package errors provides a simple way to wrap errors with a status code and a cause.
package errors

import (
	"errors"
	"fmt"

	pkgerrors "github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/saltfishpr/go-learning/pkg/util"
)

type Error struct {
	s  *status.Status
	md map[string]string

	cause error
}

func New(code codes.Code, message string) *Error {
	return &Error{
		s:     status.New(code, message),
		cause: WithStack(errors.New(message)),
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s cause = %v", e.s.String(), e.cause)
}

func (e *Error) Format(s fmt.State, verb rune) {
	const (
		skipFrames = 2
		maxFrames  = 32
	)
	st := TraceStack(e.cause)
	if len(st) <= skipFrames {
		return
	}
	st = st[skipFrames:]
	if len(st) > maxFrames {
		st = st[:maxFrames]
	}
	st.Format(s, verb)
}

// Unwrap returns the cause of the error.
func (e *Error) Unwrap() error {
	return e.cause
}

func (e *Error) Cause() error {
	return e.cause
}

func (e *Error) WithCause(cause error) *Error {
	if cause == nil {
		return e
	}
	return &Error{
		s:     e.s,
		md:    util.CloneMap(e.md),
		cause: WithStack(cause),
	}
}

func (e *Error) WithMetadataPair(key, value string) *Error {
	_md := util.CloneMap(e.md)
	if _md == nil {
		_md = make(map[string]string)
	}
	_md[key] = value
	return &Error{
		s:     e.s,
		md:    _md,
		cause: e.cause,
	}
}

// WithMetadata returns a new error with the given metadata.
func (e *Error) WithMetadata(md map[string]string) *Error {
	_md := util.CloneMap(e.md)
	if _md == nil {
		_md = make(map[string]string)
	}
	for k, v := range md {
		_md[k] = v
	}
	return &Error{
		s:     e.s,
		md:    _md,
		cause: e.cause,
	}
}

// GRPCStatus returns the gRPC status of the error.
func (e *Error) GRPCStatus() *status.Status {
	if se, ok := e.cause.(interface {
		GRPCStatus() *status.Status
	}); ok {
		return se.GRPCStatus()
	}

	if len(e.md) == 0 {
		return e.s
	}

	s, err := e.s.WithDetails(&errdetails.ErrorInfo{
		Metadata: e.md,
	})
	if err != nil {
		return e.s
	}
	return s
}

// StackTrace returns the stack trace of the error.
func (e *Error) StackTrace() pkgerrors.StackTrace {
	return TraceStack(e.cause)
}
