// @file: stock.go
// @date: 2021/11/09

package main

import "fmt"

type Stock struct {
	name     string
	quantity int
}

func (s *Stock) buy() {
	fmt.Printf("Stock: {Name: %s, Quantity: %d} bought\n", s.name, s.quantity)
}

func (s *Stock) sell() {
	fmt.Printf("Stock: {Name: %s, Quantity: %d} sold\n", s.name, s.quantity)
}
