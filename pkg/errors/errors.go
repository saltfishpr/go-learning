// Package errors provides a simple way to wrap errors with a status code and a cause.
package errors

import (
	"errors"
	"fmt"

	pkgerrors "github.com/pkg/errors"
	"github.com/samber/lo"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error struct {
	s     *status.Status
	md    map[string]string
	cause error
}

func New(code codes.Code, message string) *Error {
	return &Error{
		s:     status.New(code, message),
		cause: withStack(errors.New(message)),
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s cause = %v", e.s.String(), e.cause)
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

	if se, ok := lo.ErrorsAs[*Error](cause); ok {
		return se
	}

	return &Error{
		s:     e.s,
		cause: withStack(cause),
	}
}

func (e *Error) WithMetadataPair(key, value string) *Error {
	if e.md == nil {
		e.md = make(map[string]string)
	}
	e.md[key] = value
	return e
}

func (e *Error) WithMetadata(md map[string]string) *Error {
	if e.md == nil {
		e.md = make(map[string]string)
	}
	for k, v := range md {
		e.md[k] = v
	}
	return e
}

func (e *Error) GRPCStatus() *status.Status {
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

func Cause(err error) error {
	return pkgerrors.Cause(err)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func Is(err error, target error) bool {
	return errors.Is(err, Cause(err))
}
