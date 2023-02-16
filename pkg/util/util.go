// Package util provides some utility functions.
package util

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func CloneMap[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}
	_m := make(map[K]V, len(m))
	for k, v := range m {
		_m[k] = v
	}
	return _m
}
