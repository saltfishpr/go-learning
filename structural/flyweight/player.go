// @file: Player.go
// @date: 2021/11/9

package main

type Player struct {
	dress      Dress
	playerType string
}

func NewPlayer(playerType, dressType string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}
