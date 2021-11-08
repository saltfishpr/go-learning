// @file: printer.go
// @date: 2021/11/08

package main

import "fmt"

type Printer interface {
	PrintFile()
}

type epson struct{}

func (p *epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type hp struct{}

func (p *hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
