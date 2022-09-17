// @file: goos.go
// @description: 输出系统信息
// @author: SaltFish
// @date: 2020/07/31

// Package ch4 is chapter 4
package ch4

import (
	"fmt"
	"os"
	"runtime"
)

// Goos is a exported function
func Goos() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
