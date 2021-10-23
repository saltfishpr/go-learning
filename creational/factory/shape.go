// @file: shape.go
// @date: 2021/10/26

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
