package validator

import (
	"google.golang.org/grpc/codes"

	"github.com/saltfishpr/go-learning/pkg/errors"
)

var defaultOptions = &options{
	status: errors.NewStatus(codes.InvalidArgument, codes.InvalidArgument.String()),
}

type options struct {
	all    bool
	status *errors.StatusImpl
}

type Option func(*options)

func evaluateOptions(opts ...Option) *options {
	optCopy := &options{}
	*optCopy = *defaultOptions
	for _, opt := range opts {
		opt(optCopy)
	}
	return optCopy
}

func WithStatus(status *errors.StatusImpl) Option {
	return func(o *options) {
		o.status = status
	}
}

func WithAll(all bool) Option {
	return func(o *options) {
		o.all = all
	}
}
