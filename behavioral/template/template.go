// @description:
// @file: template.go
// @date: 2021/12/03

package main

type Game interface {
	initial()
	start()
	end()
	award()
}

type Template struct {
	Game
}

func (t Template) Play() {
	t.initial()
	t.start()
	t.end()
	t.award()
}
