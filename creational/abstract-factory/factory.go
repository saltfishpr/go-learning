// @file: factory.go
// @date: 2021/10/27

package main

type AbstractFactory interface {
	GetColor(color string) Color
	GetShape(shape string) Shape
}

type ShapeFactory struct{}

func (ShapeFactory) GetColor(string) Color {
	return nil
}

func (ShapeFactory) GetShape(shape string) Shape {
	switch shape {
	case "CIRCLE":
		return new(Circle)
	case "RECTANGLE":
		return new(Rectangle)
	case "SQUARE":
		return new(Square)
	default:
		return nil
	}
}

type ColorFactory struct{}

func (ColorFactory) GetColor(color string) Color {
	switch color {
	case "RED":
		return new(Red)
	case "GREEN":
		return new(Green)
	case "BLUE":
		return new(Blue)
	default:
		return nil
	}
}

func (ColorFactory) GetShape(string) Shape {
	return nil
}

func NewFactory(choice string) AbstractFactory {
	switch choice {
	case "SHAPE":
		return new(ShapeFactory)
	case "COLOR":
		return new(ColorFactory)
	default:
		return nil
	}
}
