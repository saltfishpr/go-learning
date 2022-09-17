// @file: main.go
// @date: 2021/10/27

package main

import (
	"fmt"
)

func main() {
	instance1 := singleton.GetInstance()
	instance2 := singleton.GetInstance()
	fmt.Println(instance1 == instance2)
}
