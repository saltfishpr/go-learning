// @file: main.go
// @date: 2021/11/9

package main

import "fmt"

func main() {
	var image Image = NewProxyImage("1.jpg")
	image.Display() // need to load from disk
	fmt.Println()
	image.Display() // no need to load from disk
}
