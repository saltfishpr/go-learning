// @file: main.go
// @date: 2021/11/09

package main

type Broker struct {
	orderList []Order
}

func (b *Broker) TakeOrder(order Order) {
	b.orderList = append(b.orderList, order)
}

func (b *Broker) PlaceOrders() {
	if b.orderList == nil {
		return
	}
	for _, order := range b.orderList {
		order.Execute()
	}
	b.orderList = nil
}

func main() {
	maotai := &Stock{name: "茅台", quantity: 2000}

	buyStock := &BuyStock{Stock: maotai}
	sellStock := &SellStock{Stock: maotai}

	broker := new(Broker)
	broker.TakeOrder(buyStock)
	broker.TakeOrder(sellStock)

	broker.PlaceOrders()
}
