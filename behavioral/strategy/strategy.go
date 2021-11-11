// @description:
// @file: strategy.go
// @date: 2021/12/03

package main

type Strategy interface {
	DoOperation(int, int) int
}

type OperationAdd struct{}

func (o OperationAdd) DoOperation(num1 int, num2 int) int {
	return num1 + num2
}

type OperationSub struct{}

func (o OperationSub) DoOperation(num1 int, num2 int) int {
	return num1 - num2
}

type OperationMul struct{}

func (o OperationMul) DoOperation(num1 int, num2 int) int {
	return num1 * num2
}
