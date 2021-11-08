// @file: game.go
// @date: 2021/11/9

package main

type Game struct {
	factory           *dressFactory
	terrorists        []*Player
	counterTerrorists []*Player
}

func NewGame() *Game {
	return &Game{factory: getDressFactorySingleInstance()}
}

func (g *Game) AddPlayer(dressType string) {
	switch dressType {
	case TerroristDressType:
		g.terrorists = append(g.terrorists, NewPlayer("T", dressType))
	case CounterTerroristDressType:
		g.terrorists = append(g.terrorists, NewPlayer("CT", dressType))
	}
}
