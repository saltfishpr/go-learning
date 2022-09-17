// @file: subject.go
// @date: 2021/11/11

package main

type Subject struct {
	observers []Observer
	state     int
}

func NewSubject() *Subject {
	return new(Subject)
}

func (s Subject) GetState() int {
	return s.state
}

func (s *Subject) SetState(state int) {
	s.state = state
	s.notifyAllObservers()
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s Subject) notifyAllObservers() {
	for _, observer := range s.observers {
		observer.Update()
	}
}
