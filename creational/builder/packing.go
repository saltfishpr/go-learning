// @file: Packing.go
// @date: 2021/10/27

package main

type Packing interface {
	Pack() string
}

type Wrapper struct{}

func (Wrapper) Pack() string {
	return "wrapper"
}

type Bottle struct{}

func (Bottle) Pack() string {
	return "bottle"
}
