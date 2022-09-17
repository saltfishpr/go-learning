// @file: shape.go
// @date: 2021/10/27
// @describe：

package main

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct{}

func (Circle) Draw() {
	fmt.Println("Circle Draw().")
}

type Rectangle struct{}

func (Rectangle) Draw() {
	fmt.Println("Rectangle Draw().")
}

type Square struct{}

func (Square) Draw() {
	fmt.Println("Square Draw().")
}
