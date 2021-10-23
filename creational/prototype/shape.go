// @file: shape.go
// @date: 2021/10/28

package main

import "fmt"

type Cloneable interface {
	clone() Shape
}

type Drawable interface {
	draw()
}

type Shape interface {
	Cloneable
	Drawable

	getId() string
	setId(string)
	getTp() string
}

type ShapeBase struct {
	id string
	tp string
}

func (s ShapeBase) getId() string {
	return s.id
}

func (s *ShapeBase) setId(id string) {
	s.id = id
}

func (s ShapeBase) getTp() string {
	return s.tp
}

type Circle struct {
	*ShapeBase
}

func NewCircle() *Circle {
	return &Circle{&ShapeBase{tp: "Circle"}}
}

func (c Circle) clone() Shape {
	return &Circle{&ShapeBase{id: c.id, tp: c.tp}}
}

func (Circle) draw() {
	fmt.Println("Circle draw().")
}

type Rectangle struct {
	*ShapeBase
}

func NewRectangle() *Rectangle {
	return &Rectangle{&ShapeBase{tp: "Rectangle"}}
}

func (r Rectangle) clone() Shape {
	return &Rectangle{&ShapeBase{id: r.id, tp: r.tp}}
}

func (Rectangle) draw() {
	fmt.Println("Rectangle draw().")
}

type Square struct {
	*ShapeBase
}

func NewSquare() *Square {
	return &Square{&ShapeBase{tp: "Square"}}
}

func (s Square) clone() Shape {
	return &Square{&ShapeBase{id: s.id, tp: s.tp}}
}

func (Square) draw() {
	fmt.Println("Square draw().")
}
