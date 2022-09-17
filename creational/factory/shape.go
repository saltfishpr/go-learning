// @file: shape.go
// @date: 2021/10/26

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

// NewShape return a Shape decide by shape
func NewShape(shape string) Shape {
	switch shape {
	case "CIRCLE":
		return new(Circle)
	case "RECTANGLE":
		return new(Rectangle)
	case "SQUARE":
		return new(Square)
	default:
		return nil
	}
}
