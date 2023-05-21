// Package util provides some utility functions.
package util

import (
	"context"
	"os"
	"os/signal"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func CloneSlice[S ~[]T, T any](s S) S {
	if s == nil {
		return nil
	}
	_s := make(S, len(s))
	copy(_s, s)
	return _s
}

func CloneMap[M ~map[K]V, K comparable, V any](m M) M {
	if m == nil {
		return nil
	}
	_m := make(M, len(m))
	for k, v := range m {
		_m[k] = v
	}
	return _m
}

// ContextForSignal returns a context object which is cancelled when a signal
// is received. It returns nil if no signal parameter is provided
func ContextForSignal(signals ...os.Signal) context.Context {
	if len(signals) == 0 {
		return nil
	}

	ch := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())

	// Send message on channel when signal received
	signal.Notify(ch, signals...)

	// When any signal received, call cancel
	go func() {
		<-ch
		cancel()
	}()

	// Return success
	return ctx
}
