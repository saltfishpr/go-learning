// @file: doctor.go
// @date: 2021/11/9

package main

import "fmt"

type Doctor struct {
	next ChainNode
}

func (d *Doctor) Execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("doctor checkup already done")
		d.next.Execute(p)
		return
	}
	fmt.Println("doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.Execute(p)
}

func (d *Doctor) SetNext(next ChainNode) {
	d.next = next
}
