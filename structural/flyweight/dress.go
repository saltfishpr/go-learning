// @file: dress.go
// @date: 2021/11/9

package main

type Dress interface {
	GetColor() string
}

type terroristDress struct {
	color string
}

func (t *terroristDress) GetColor() string {
	return t.color
}

func newTerroristDress() *terroristDress {
	return &terroristDress{color: "red"}
}

type counterTerroristDress struct {
	color string
}

func (c *counterTerroristDress) GetColor() string {
	return c.color
}

func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{color: "green"}
}
