// @file: order.go
// @date: 2021/11/09

package main

type Order interface {
	Execute()
}

type BuyStock struct {
	*Stock
}

func (s BuyStock) Execute() {
	s.buy()
}

type SellStock struct {
	*Stock
}

func (s SellStock) Execute() {
	s.sell()
}
