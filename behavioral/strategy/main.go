// @description:
// @file: main.go
// @date: 2021/12/03

package main

import "fmt"

type Context struct {
	strategy Strategy
}

func (c Context) exec(num1 int, num2 int) int {
	return c.strategy.DoOperation(num1, num2)
}

func NewContext(strategy Strategy) *Context {
	return &Context{strategy: strategy}
}

func main() {
	var ctx *Context
	ctx = NewContext(OperationAdd{})
	fmt.Println(ctx.exec(5, 2))

	ctx = NewContext(OperationSub{})
	fmt.Println(ctx.exec(5, 2))

	ctx = NewContext(OperationMul{})
	fmt.Println(ctx.exec(5, 2))
}
