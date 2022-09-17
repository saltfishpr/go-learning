// @file: main.go
// @date: 2021/10/27

package main

import "fmt"

func main() {
	shapeCache := NewShapeCache()
	shapeCache.loadCache()

	clonedShape1 := shapeCache.getShape("1")
	fmt.Println("Shape: ", clonedShape1.GetTp())
	clonedShape1.Draw()

	clonedShape2 := shapeCache.getShape("2")
	fmt.Println("Shape: ", clonedShape2.GetTp())
	clonedShape2.Draw()

	clonedShape3 := shapeCache.getShape("3")
	fmt.Println("Shape: ", clonedShape3.GetTp())
	clonedShape3.Draw()

	clonedShape4 := shapeCache.getShape("4")
	_ = clonedShape4 == nil
}
