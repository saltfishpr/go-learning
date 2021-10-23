// @file: burger.go
// @date: 2021/10/27

package main

type Burger struct{}

func (Burger) packing() Packing {
	return new(Wrapper)
}

type VegBurger struct {
	Burger
}

func (VegBurger) name() string {
	return "Veg Burger"
}

func (VegBurger) price() int {
	return 25
}

type ChickenBurger struct {
	Burger
}

func (ChickenBurger) name() string {
	return "Chicken Burger"
}

func (ChickenBurger) price() int {
	return 40
}
