// @file: cold_drink.go
// @date: 2021/10/27

package main

type ColdDrink struct{}

func (ColdDrink) Packing() Packing {
	return new(Bottle)
}

type Coke struct {
	ColdDrink
}

func (Coke) Name() string {
	return "Coke"
}

func (Coke) Price() int {
	return 10
}

type Pepsi struct {
	ColdDrink
}

func (Pepsi) Name() string {
	return "Pepsi"
}

func (Pepsi) Price() int {
	return 15
}
