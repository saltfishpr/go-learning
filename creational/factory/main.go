// @file: main.go
// @date: 2021/10/26

package main

func main() {
	shape1 := NewShape("CIRCLE")
	if shape1 != nil {
		shape1.draw()
	}

	shape2 := NewShape("RECTANGLE")
	if shape2 != nil {
		shape2.draw()
	}

	shape3 := NewShape("SQUARE")
	if shape3 != nil {
		shape3.draw()
	}

	shape4 := NewShape("") // nil
	if shape4 != nil {
		shape4.draw()
	}
}
