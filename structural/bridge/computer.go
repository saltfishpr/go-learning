// @file: computer.go
// @date: 2021/11/08

package main

import "fmt"

type Computer interface {
	Print()
	SetPrinter(Printer)
}

type windows struct {
	printer Printer
}

func (w *windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *windows) SetPrinter(p Printer) {
	w.printer = p
}

type mac struct {
	printer Printer
}

func (m *mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *mac) SetPrinter(p Printer) {
	m.printer = p
}
