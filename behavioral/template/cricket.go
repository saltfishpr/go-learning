// @description:
// @file: cricket.go
// @date: 2021/12/03

package main

import "fmt"

type Cricket struct{}

func (c Cricket) initial() {
	fmt.Println("Cricket initial")
}

func (c Cricket) start() {
	fmt.Println("Cricket game start")
}

func (c Cricket) end() {
	fmt.Println("Cricket game end")
}

func (c Cricket) award() {
	fmt.Println("Cricket award")
}
