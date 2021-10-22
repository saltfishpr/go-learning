// @file: gotemplate.go
// @description: Go语言基础结构模板
// @author: SaltFish
// @date: 2020/07/31

// Package ch4 is chapter 4
package ch4

import (
	"fmt"
)

// constant values
const c = "C"

// variable values
var v = 5

// T is a struct template
type T struct{}

// initialization of package
func init() {
}

// GoTemplate is a exported function template
func GoTemplate() {
	var a int
	a = v
	Func1()
	// ...
	fmt.Println(a)
}

// Method1 is a exported function of type T
func (t T) Method1() {
	// ...
}

// Func1 is a exported function of package
func Func1() {
	// ...
}
