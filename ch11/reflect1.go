// @file: reflect1.go
// @description:
// @author: SaltFish
// @date: 2020/08/19

// Package ch11 is chapter 11
package ch11

import (
	"fmt"
	"reflect"
)

func Reflect1() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}
