// @file: cashier.go
// @date: 2021/11/9

package main

import "fmt"

type Cashier struct {
	next ChainNode
}

func (c *Cashier) Execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("payment done")
	}
	fmt.Println("cashier getting money from patient patient")
}

func (c *Cashier) SetNext(next ChainNode) {
	c.next = next
}
