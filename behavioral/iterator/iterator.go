// @file: iterator.go
// @date: 2021/11/09

package main

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type NameIterator struct {
	index int
	names []string
}

func NewNameIterator(names []string) *NameIterator {
	return &NameIterator{names: names}
}

func (r NameIterator) HasNext() bool {
	return r.index < len(r.names)
}

func (r *NameIterator) Next() interface{} {
	if r.HasNext() {
		res := r.names[r.index]
		r.index++
		return res
	}
	return nil
}
