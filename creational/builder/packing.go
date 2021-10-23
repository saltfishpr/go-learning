// @file: packing.go
// @date: 2021/10/27

package main

type Packing interface {
	pack() string
}

type Wrapper struct{}

func (Wrapper) pack() string {
	return "wrapper"
}

type Bottle struct{}

func (Bottle) pack() string {
	return "bottle"
}
