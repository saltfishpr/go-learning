// @file: main.go
// @date: 2021/11/08

package main

func main() {
	var bridge Computer
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	bridge = new(mac)
	bridge.SetPrinter(hpPrinter)
	bridge.Print()
	bridge.SetPrinter(epsonPrinter)
	bridge.Print()

	bridge = new(windows)
	bridge.SetPrinter(hpPrinter)
	bridge.Print()
	bridge.SetPrinter(epsonPrinter)
	bridge.Print()
}
