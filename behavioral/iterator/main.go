// @file: main.go
// @date: 2021/11/09

package main

import "fmt"

func main() {
	nameIterator := NewNameIterator([]string{"Robert", "John", "Julie", "Lora"})
	for nameIterator.HasNext() {
		fmt.Println("Name: ", nameIterator.Next())
	}
}
