// @description:
// @file: shape.go
// @date: 2021/12/03

package main

type Shape interface {
	GetType() string
	Accept(visitor)
}

type Square struct {
	side int
}

func (s *Square) Accept(v visitor) {
	v.visit(s)
}

func (s *Square) GetType() string {
	return "Square"
}

type Circle struct {
	radius int
}

func (c *Circle) Accept(v visitor) {
	v.visit(c)
}

func (c *Circle) GetType() string {
	return "Circle"
}

type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) Accept(v visitor) {
	v.visit(t)
}

func (t *Rectangle) GetType() string {
	return "rectangle"
}
