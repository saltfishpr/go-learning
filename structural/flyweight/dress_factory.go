// @file: dress_factory.go
// @date: 2021/11/9

package main

import "fmt"

const (
	// TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	// CounterTerroristDressType terrorist dress type
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]Dress),
	}
)

type dressFactory struct {
	dressMap map[string]Dress
}

func (d *dressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("wrong dress type passed")
}

func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}
