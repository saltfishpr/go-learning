// @file: shape.go
// @date: 2021/10/28

package main

import "fmt"

type Cloneable interface {
	Clone() Shape
}

type Drawable interface {
	Draw()
}

type Shape interface {
	Cloneable
	Drawable

	GetId() string
	SetId(string)
	GetTp() string
}

type ShapeBase struct {
	id string
	tp string
}

func (s ShapeBase) GetId() string {
	return s.id
}

func (s *ShapeBase) SetId(id string) {
	s.id = id
}

func (s ShapeBase) GetTp() string {
	return s.tp
}

type Circle struct {
	*ShapeBase
}

func NewCircle() *Circle {
	return &Circle{&ShapeBase{tp: "Circle"}}
}

func (c Circle) Clone() Shape {
	return &Circle{&ShapeBase{id: c.id, tp: c.tp}}
}

func (Circle) Draw() {
	fmt.Println("Circle Draw().")
}

type Rectangle struct {
	*ShapeBase
}

func NewRectangle() *Rectangle {
	return &Rectangle{&ShapeBase{tp: "Rectangle"}}
}

func (r Rectangle) Clone() Shape {
	return &Rectangle{&ShapeBase{id: r.id, tp: r.tp}}
}

func (Rectangle) Draw() {
	fmt.Println("Rectangle Draw().")
}

type Square struct {
	*ShapeBase
}

func NewSquare() *Square {
	return &Square{&ShapeBase{tp: "Square"}}
}

func (s Square) Clone() Shape {
	return &Square{&ShapeBase{id: s.id, tp: s.tp}}
}

func (Square) Draw() {
	fmt.Println("Square Draw().")
}
