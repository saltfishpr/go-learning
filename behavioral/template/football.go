// @description:
// @file: football.go
// @date: 2021/12/03

package main

import "fmt"

type Football struct{}

func (f Football) initial() {
	fmt.Println("Football initial")
}

func (f Football) start() {
	fmt.Println("Football game start")
}

func (f Football) end() {
	fmt.Println("Football game end")
}

func (f Football) award() {
	fmt.Println("Football award")
}
