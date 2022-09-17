// @file: main.go
// @date: 2021/11/8

package main

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct{}

func (Circle) Draw() {
	fmt.Println("Shape: Circle.")
}

type Rectangle struct{}

func (Rectangle) Draw() {
	fmt.Println("Shape: Rectangle.")
}
