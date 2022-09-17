// @file: main.go
// @date: 2021/11/9

package main

import "fmt"

func main() {
	game := NewGame()

	game.AddPlayer(TerroristDressType)
	game.AddPlayer(TerroristDressType)
	game.AddPlayer(TerroristDressType)
	game.AddPlayer(TerroristDressType)
	game.AddPlayer(TerroristDressType)

	game.AddPlayer(CounterTerroristDressType)
	game.AddPlayer(CounterTerroristDressType)
	game.AddPlayer(CounterTerroristDressType)
	game.AddPlayer(CounterTerroristDressType)
	game.AddPlayer(CounterTerroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()
	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}
