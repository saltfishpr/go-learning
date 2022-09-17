// @description:
// @file: main.go
// @date: 2021/12/03

package main

import "fmt"

type Customer interface {
	IsNil() bool
	GetName() string
}

type RealCustomer struct {
	name string
}

func (c RealCustomer) GetName() string {
	return c.name
}

func (c RealCustomer) IsNil() bool {
	return false
}

type NullCustomer struct{}

func (c NullCustomer) GetName() string {
	return "Not Available in Customer Database"
}

func (c NullCustomer) IsNil() bool {
	return true
}

func NewCustomer(name string) Customer {
	names := []string{"Alice", "Bob"}
	for i := range names {
		if names[i] == name {
			return RealCustomer{name: name}
		}
	}
	return NullCustomer{}
}

func main() {
	customer1 := NewCustomer("Alice")
	customer2 := NewCustomer("Bob")
	customer3 := NewCustomer("WWS")

	fmt.Println(customer1.GetName())
	fmt.Println(customer2.GetName())
	fmt.Println(customer3.GetName())
}
