// @file: user_init.go
// @description:
// @author: SaltFish
// @date: 2020/07/31

// Package ch4 is chapter 4
package ch4

import (
	"fmt"

	"learning/ch4/trans"
)

var twoPi = 2 * trans.Pi

// TwoPi can get 2 * pi
func TwoPi() {
	fmt.Printf("2*Pi = %g\n", twoPi) // 2*Pi = 6.283185307179586
}
