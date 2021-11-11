// @file: expression.go
// @date: 2021/11/09

package main

import "strings"

type Expression interface {
	Interpret(text string) bool
}

type TerminalExpression struct {
	data string
}

func (e TerminalExpression) Interpret(text string) bool {
	return strings.Contains(text, e.data)
}

type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func (e OrExpression) Interpret(text string) bool {
	return e.expr1.Interpret(text) || e.expr2.Interpret(text)
}

type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func (e AndExpression) Interpret(text string) bool {
	return e.expr1.Interpret(text) && e.expr2.Interpret(text)
}
