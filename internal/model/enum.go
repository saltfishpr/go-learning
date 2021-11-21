// @description: 一些枚举类型
// @file: enum.go
// @date: 2021/11/22

package model

type ChatMode int

const (
	ToUser ChatMode = iota
	ToHub
)
