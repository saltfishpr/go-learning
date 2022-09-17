// @file: reception.go
// @date: 2021/11/9

package main

import "fmt"

type Reception struct {
	next ChainNode
}

func (r *Reception) Execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("patient registration already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("reception registering patient")
	p.registrationDone = true
	r.next.Execute(p)
}

func (r *Reception) SetNext(next ChainNode) {
	r.next = next
}
