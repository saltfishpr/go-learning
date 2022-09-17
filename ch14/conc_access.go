// @file: conc_access.go
// @description:
// @author: SaltFish
// @date: 2020/09/08

// Package ch14 is chapter 14
package ch14

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name   string
	salary float64
	chF    chan func()
}

func NewPerson(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend()
	return p
}

// backend()方法会在一个无限循环中执行chF中放置的所有函数，有效的将它们序列化从而提供了安全的并发访问。
// 更改和读取salary的方法会通过将一个匿名函数写入chF通道中，然后让backend()按顺序执行以达到其目的。
func (p *Person) backend() {
	for f := range p.chF {
		f()
	}
}

// Set salary.
func (p *Person) SetSalary(sal float64) {
	p.chF <- func() { p.salary = sal }
}

// Retrieve salary.
func (p *Person) Salary() float64 {
	fChan := make(chan float64)
	p.chF <- func() { fChan <- p.salary }
	return <-fChan
}

func (p *Person) String() string {
	return "Person - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(
		p.Salary(),
		'f',
		2,
		64,
	)
}

func ConAccess() {
	bs := NewPerson("Smith Bill", 2500.5)
	fmt.Println(bs)
	bs.SetSalary(4000.25)
	fmt.Println("Salary changed:")
	fmt.Println(bs)
}
