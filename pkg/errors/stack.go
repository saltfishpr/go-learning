package errors

import (
	pkgerrors "github.com/pkg/errors"
)

type StackTracer interface {
	StackTrace() pkgerrors.StackTrace
}

func WithStack(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(StackTracer); ok {
		return err
	}
	return pkgerrors.WithStack(err)
}

func HasStack(err error) bool {
	if err == nil {
		return false
	}
	if _, ok := err.(StackTracer); ok {
		return true
	}
	return false
}

func TraceStack(err error) pkgerrors.StackTrace {
	if err == nil {
		return nil
	}
	if st, ok := err.(StackTracer); ok {
		return st.StackTrace()
	}
	return nil
}
