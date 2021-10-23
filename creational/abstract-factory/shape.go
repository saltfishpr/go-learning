// @file: shape.go
// @date: 2021/10/27
// @describe：

package main

import "fmt"

type Shape interface {
	draw()
}

type Circle struct{}

func (Circle) draw() {
	fmt.Println("Circle draw().")
}

type Rectangle struct{}

func (Rectangle) draw() {
	fmt.Println("Rectangle draw().")
}

type Square struct{}

func (Square) draw() {
	fmt.Println("Square draw().")
}
