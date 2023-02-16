package errors

import (
	pkgerrors "github.com/pkg/errors"
	"github.com/samber/lo"
)

func WithStack(err error) error {
	if err == nil {
		return nil
	}

	if se, ok := lo.ErrorsAs[*Error](err); ok {
		return se
	}

	return withStack(err)
}

type stackTracer interface {
	StackTrace() pkgerrors.StackTrace
}

func withStack(err error) error {
	if err == nil {
		return nil
	}

	if _, ok := err.(stackTracer); ok {
		return err
	}

	return pkgerrors.WithStack(err)
}

func StackTrace(err error) pkgerrors.StackTrace {
	if err == nil {
		return nil
	}

	if st, ok := err.(stackTracer); ok {
		return st.StackTrace()
	}

	return nil
}
