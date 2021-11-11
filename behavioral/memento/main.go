// @file: main.go
// @date: 2021/11/11

package main

import "fmt"

type Memento struct {
	state string
}

func (m Memento) GetState() string {
	return m.state
}

type Originator struct {
	State string
}

func (o Originator) SaveStateToMemento() *Memento {
	return &Memento{state: o.State}
}

func (o *Originator) GetStateFromMemento(memento Memento) {
	o.State = memento.GetState()
}

type CareTaker struct {
	mementoList []*Memento
}

func (c *CareTaker) Add(memento *Memento) {
	c.mementoList = append(c.mementoList, memento)
}

func (c CareTaker) Get(index int) *Memento {
	return c.mementoList[index]
}

func main() {
	originator := new(Originator)
	careTaker := new(CareTaker)
	originator.State = "State #1"
	originator.State = "State #2"
	careTaker.Add(originator.SaveStateToMemento())
	originator.State = "State #3"
	careTaker.Add(originator.SaveStateToMemento())
	originator.State = "State #4"

	fmt.Println("Current state: ", originator.State)
	originator.GetStateFromMemento(*careTaker.Get(0))
	fmt.Println("First saved state: ", originator.State)
	originator.GetStateFromMemento(*careTaker.Get(1))
	fmt.Println("Second saved state: ", originator.State)
}
