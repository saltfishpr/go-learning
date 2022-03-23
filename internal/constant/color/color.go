// @file: color.go
// @date: 2021/11/25

// Package color 终端彩色输出的控制字符.
package color

type Color string

const (
	Reset       Color = "\033[0m"
	Red         Color = "\033[31m"
	Green       Color = "\033[32m"
	Yellow      Color = "\033[33m"
	Blue        Color = "\033[34m"
	Magenta     Color = "\033[35m"
	Cyan        Color = "\033[36m"
	White       Color = "\033[37m"
	BlueBold    Color = "\033[34;1m"
	MagentaBold Color = "\033[35;1m"
	RedBold     Color = "\033[31;1m"
	YellowBold  Color = "\033[33;1m"
)
