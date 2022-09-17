// @description:
// @file: main.go
// @date: 2021/12/03

package main

import "fmt"

func main() {
	var game *Template

	game = &Template{Game: Cricket{}}
	game.Play()

	fmt.Println()

	game = &Template{Game: Football{}}
	game.Play()
}
