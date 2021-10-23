// @file: singleton.go
// @date: 2021/10/27

package main

import "sync"

var singleton = Singleton{}

type Instance struct{}

type Singleton struct {
	instance *Instance
	once     sync.Once
}

func (s *Singleton) GetInstance() *Instance {
	s.once.Do(
		func() {
			s.instance = new(Instance)
		},
	)

	return s.instance
}
