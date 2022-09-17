// @file: builder.go
// @date: 2021/10/27

package main

import "fmt"

type Item interface {
	Name() string
	Packing() Packing
	Price() int
}

type Meal struct {
	items []Item
}

func (m *Meal) AddItem(item Item) {
	m.items = append(m.items, item)
}

func (m Meal) GetCost() int {
	var res int
	for _, item := range m.items {
		res += item.Price()
	}
	return res
}

func (m Meal) showItems() {
	for _, item := range m.items {
		fmt.Printf("Item: %s, Packing: %s, Price: %d\n", item.Name(), item.Packing().Pack(), item.Price())
	}
}

type MealBuilder struct{}

func (MealBuilder) PrepareVegMeal() *Meal {
	meal := new(Meal)
	meal.AddItem(new(VegBurger))
	meal.AddItem(new(Coke))
	return meal
}

func (MealBuilder) PrepareNonVegMeal() *Meal {
	meal := new(Meal)
	meal.AddItem(new(ChickenBurger))
	meal.AddItem(new(Pepsi))
	return meal
}
