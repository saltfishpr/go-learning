// @file: medical.go
// @date: 2021/11/9

package main

import "fmt"

type Medical struct {
	next ChainNode
}

func (m *Medical) Execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("medicine already given to patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("medical giving medicine to patient")
	p.medicineDone = true
	m.next.Execute(p)
}

func (m *Medical) SetNext(next ChainNode) {
	m.next = next
}
