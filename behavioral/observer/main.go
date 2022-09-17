// @file: main.go
// @date: 2021/11/11

package main

import "fmt"

func main() {
	subject := NewSubject()

	NewHexaObserver(subject)
	NewOctalObserver(subject)
	NewBinaryObserver(subject)

	fmt.Println("state change: 15")
	subject.SetState(15)
	fmt.Println()
	fmt.Println("state change: 10")
	subject.SetState(10)
}
