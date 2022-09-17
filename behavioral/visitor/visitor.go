// @description:
// @file: visitor.go
// @date: 2021/12/03

package main

import "fmt"

type visitor interface {
	visit(Shape)
}

type AreaCalculator struct{}

func (a *AreaCalculator) visit(s Shape) {
	fmt.Println("Calculating area for", s.GetType())
}

type MiddleCoordinates struct{}

func (a *MiddleCoordinates) visit(s Shape) {
	fmt.Println("Calculating middle point coordinates for", s.GetType())
}
