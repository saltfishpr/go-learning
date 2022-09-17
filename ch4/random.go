// @file: random.go
// @description:
// @author: SaltFish
// @date: 2020/07/31

// Package ch4 is chapter 4
package ch4

import (
	"fmt"
	"math/rand"
	"time"
)

// MyRandom is fun
func MyRandom() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d / ", r)
	}
	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}
}
