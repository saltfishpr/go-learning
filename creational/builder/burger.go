// @file: burger.go
// @date: 2021/10/27

package main

type Burger struct{}

func (Burger) Packing() Packing {
	return new(Wrapper)
}

type VegBurger struct {
	Burger
}

func (VegBurger) Name() string {
	return "Veg Burger"
}

func (VegBurger) Price() int {
	return 25
}

type ChickenBurger struct {
	Burger
}

func (ChickenBurger) Name() string {
	return "Chicken Burger"
}

func (ChickenBurger) Price() int {
	return 40
}
