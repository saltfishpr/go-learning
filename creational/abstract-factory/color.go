// @file: color.go
// @date: 2021/10/27

package main

import "fmt"

type Color interface {
	Fill()
}

type Red struct{}

func (Red) Fill() {
	fmt.Println("Red Fill().")
}

type Green struct{}

func (Green) Fill() {
	fmt.Println("Green Fill().")
}

type Blue struct{}

func (Blue) Fill() {
	fmt.Println("Blue Fill().")
}
