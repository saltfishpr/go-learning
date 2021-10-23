// @file: main.go
// @date: 2021/10/27

package main

import "fmt"

func main() {
	mealBuilder := new(MealBuilder)

	vegMeal := mealBuilder.PrepareVegMeal()
	vegMeal.showItems()
	fmt.Println("Total cost: ", vegMeal.GetCost())
	fmt.Println()

	nonVegMeal := mealBuilder.PrepareNonVegMeal()
	nonVegMeal.showItems()
	fmt.Println("Total cost: ", nonVegMeal.GetCost())
}
