// @file: decorator.go
// @date: 2021/11/9

package main

import "fmt"

type RedShapeDecorator struct {
	decoratedShape Shape
}

func NewRedShapeDecorator(shape Shape) *RedShapeDecorator {
	return &RedShapeDecorator{decoratedShape: shape}
}

func (r RedShapeDecorator) Draw() {
	r.decoratedShape.Draw()
	r.setRedBorder()
}

func (r RedShapeDecorator) setRedBorder() {
	fmt.Println("Border Color: Red.")
}
