// @file: my_pointer.go
// @description:
// @author: SaltFish
// @date: 2020/08/01

// Package ch4 is chapter 4
package ch4

import "fmt"

// MyPointer is fun
func MyPointer() {
	fun1()
}

func fun1() {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}
