package selector

import (
	"context"
	"strings"
)

var defaultOptions = &options{
	matchFunc: func(ctx context.Context, name string) bool {
		return false
	},
}

type options struct {
	matchFunc MatchFunc
	prefixes  []string
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

func (o *options) match(ctx context.Context, name string) bool {
	if o.matchFunc(ctx, name) {
		return true
	}
	for _, prefix := range o.prefixes {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}
	return false
}

func WithMatchFunc(f MatchFunc) Option {
	return func(o *options) {
		o.matchFunc = f
	}
}

func AddPrefix(prefix string) Option {
	return func(o *options) {
		o.prefixes = append(o.prefixes, prefix)
	}
}
