// @file: main.go
// @date: 2021/11/9

package main

import "fmt"

func main() {
	circle := new(Circle)
	redCircle := NewRedShapeDecorator(new(Circle))
	redRectangle := NewRedShapeDecorator(new(Rectangle))

	fmt.Println("Circle with normal border.")
	circle.Draw()
	fmt.Println("Circle of red border.")
	redCircle.Draw()
	fmt.Println("Rectangle of red border.")
	redRectangle.Draw()
}
