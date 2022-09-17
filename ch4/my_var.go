// @file: function_calls_function.go
// @description:
// @author: SaltFish
// @date: 2020/07/31

// Package ch4 is chapter 4
package ch4

var a = "G"
var b string

func funcCallsFunc() {
	b = "G"
	print(b)
	f1()
}

func f1() {
	b := "O"
	print(b)
	f2()
}

func f2() {
	print(b)
}

func localScope() {
	n2()
	m2()
	n2()
}

func n2() { print(a) }

func m2() {
	a := "O"
	print(a)
}

// MyVar is fun
func MyVar() {
	localScope()
}
