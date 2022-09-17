// @file: main.go
// @date: 2021/11/09

package main

import "fmt"

func main() {
	isMale := getMaleExpression()
	isMarriedWoman := getMarriedWomanExpression()

	fmt.Println("John is male? ", isMale.Interpret("John"))
	fmt.Println("Julie is a married women? ", isMarriedWoman.Interpret("Married Julie"))
}

func getMaleExpression() Expression {
	robert := TerminalExpression{data: "Robert"}
	john := TerminalExpression{data: "John"}
	return OrExpression{robert, john}
}

func getMarriedWomanExpression() Expression {
	julie := TerminalExpression{data: "Julie"}
	married := TerminalExpression{data: "Married"}
	return AndExpression{julie, married}
}
