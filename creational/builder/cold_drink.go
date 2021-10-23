// @file: cold_drink.go
// @date: 2021/10/27

package main

type ColdDrink struct{}

func (ColdDrink) packing() Packing {
	return new(Bottle)
}

type Coke struct {
	ColdDrink
}

func (Coke) name() string {
	return "Coke"
}

func (Coke) price() int {
	return 10
}

type Pepsi struct {
	ColdDrink
}

func (Pepsi) name() string {
	return "Pepsi"
}

func (Pepsi) price() int {
	return 15
}
