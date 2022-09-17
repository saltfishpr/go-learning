// @file: main.go
// @date: 2021/11/11

package main

import (
	"fmt"
)

type Context struct {
	state State
}

type State interface {
	DoAction(*Context)
}

type StartState struct{}

func (s StartState) DoAction(ctx *Context) {
	fmt.Println("Player is in start state")
	ctx.state = s
}

func (s StartState) String() string {
	return "Start State"
}

type StopState struct{}

func (s StopState) DoAction(ctx *Context) {
	fmt.Println("Player is in stop state")
	ctx.state = s
}

func (s StopState) String() string {
	return "Stop State"
}

func main() {
	context := new(Context)

	startState := new(StartState)
	startState.DoAction(context)
	fmt.Println(context.state)

	stopState := new(StopState)
	stopState.DoAction(context)
	fmt.Println(context.state)
}
