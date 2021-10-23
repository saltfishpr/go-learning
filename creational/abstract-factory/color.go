// @file: color.go
// @date: 2021/10/27

package main

import "fmt"

type Color interface {
	fill()
}

type Red struct{}

func (Red) fill() {
	fmt.Println("Red fill().")
}

type Green struct{}

func (Green) fill() {
	fmt.Println("Green fill().")
}

type Blue struct{}

func (Blue) fill() {
	fmt.Println("Blue fill().")
}
