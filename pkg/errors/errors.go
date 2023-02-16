// Package errors provides a simple way to wrap errors with a status code and a cause.
package errors

import (
	"fmt"

	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Status interface {
	fmt.Stringer
	GRPCStatus() *status.Status
}

type Error struct {
	s     Status
	cause error
}

func New(s Status) *Error {
	if s == nil {
		return nil
	}
	return &Error{s: s}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s cause = %v", e.s.String(), e.cause)
}

// Unwrap returns the cause of the error.
func (e *Error) Unwrap() error {
	return e.cause
}

func (e *Error) WithCause(cause error) error {
	if cause == nil {
		return e
	}

	if se, ok := lo.ErrorsAs[*Error](cause); ok {
		return se
	}

	return &Error{
		s:     e.s,
		cause: withStack(cause),
	}
}

func (e *Error) GRPCStatus() *status.Status {
	return e.s.GRPCStatus()
}

func FromError(err error) *Error {
	if err == nil {
		return nil
	}

	if se, ok := lo.ErrorsAs[*Error](err); ok {
		return se
	}

	if gs, ok := status.FromError(err); ok {
		return &Error{
			s:     NewStatus(gs.Code(), gs.Message()),
			cause: withStack(gs.Err()),
		}
	}

	return &Error{
		s:     NewStatus(codes.Unknown, codes.Unknown.String()),
		cause: withStack(err),
	}
}
