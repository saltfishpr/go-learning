// @file: main.go
// @date: 2021/11/08

package main

import "fmt"

type Employee struct {
	name         string
	dept         string
	salary       int
	subordinates map[*Employee]struct{}
}

func NewEmployee(name string, dept string, salary int) *Employee {
	return &Employee{name: name, dept: dept, salary: salary, subordinates: make(map[*Employee]struct{})}
}

func (e *Employee) add(employee *Employee) {
	if _, ok := e.subordinates[employee]; ok {
		return
	}
	e.subordinates[employee] = struct{}{}
}

func (e *Employee) remove(employee *Employee) {
	if _, ok := e.subordinates[employee]; !ok {
		return
	}
	delete(e.subordinates, employee)
}

func (e Employee) String() string {
	return fmt.Sprintf("Employee: {Name: %s, dept: %s, salary: %d}", e.name, e.dept, e.salary)
}

func main() {
	CEO := NewEmployee("John", "CEO", 30000)
	headSales := NewEmployee("Robert", "Head Sales", 20000)
	headMarketing := NewEmployee("Michel", "Head Marketing", 20000)
	clerk1 := NewEmployee("Laura", "Marketing", 10000)
	clerk2 := NewEmployee("Bob", "Marketing", 10000)
	salesExecutive1 := NewEmployee("Richard", "Sales", 10000)
	salesExecutive2 := NewEmployee("Rob", "Sales", 10000)

	CEO.add(headSales)
	CEO.add(headMarketing)
	headSales.add(salesExecutive1)
	headSales.add(salesExecutive2)
	headMarketing.add(clerk1)
	headMarketing.add(clerk2)

	fmt.Println(CEO)
	for e := range CEO.subordinates {
		fmt.Println(e)
		for ee := range e.subordinates {
			fmt.Println(ee)
		}
	}
}
