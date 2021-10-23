// @file: main.go
// @date: 2021/10/27

package main

func main() {
	shapeFactory := NewFactory("SHAPE")
	if shapeFactory != nil {
		shape1 := shapeFactory.GetShape("CIRCLE")
		if shape1 != nil {
			shape1.draw()
		}

		shape2 := shapeFactory.GetShape("RECTANGLE")
		if shape2 != nil {
			shape2.draw()
		}

		shape3 := shapeFactory.GetShape("SQUARE")
		if shape3 != nil {
			shape3.draw()
		}

		shape4 := shapeFactory.GetShape("") // nil
		if shape4 != nil {
			shape4.draw()
		}
	}

	colorFactory := NewFactory("COLOR")
	if colorFactory != nil {
		color1 := colorFactory.GetColor("RED")
		if color1 != nil {
			color1.fill()
		}

		color2 := colorFactory.GetColor("GREEN")
		if color2 != nil {
			color2.fill()
		}

		color3 := colorFactory.GetColor("BLUE")
		if color3 != nil {
			color3.fill()
		}

		color4 := colorFactory.GetColor("BLACK") // nil
		if color4 != nil {
			color4.fill()
		}
	}

	nilFactory := NewFactory("PRICE") // nil
	if nilFactory != nil {
		shape := shapeFactory.GetShape("CIRCLE")
		if shape != nil {
			shape.draw()
		}

		color := colorFactory.GetColor("RED")
		if color != nil {
			color.fill()
		}
	}
}
