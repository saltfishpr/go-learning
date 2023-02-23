package validator

var defaultOptions = &options{}

type options struct {
	all bool
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

func WithAll(all bool) Option {
	return func(o *options) {
		o.all = all
	}
}
