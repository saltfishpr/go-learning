// @file: ncpu_process.go
// @description: 在多核心上并行计算,信号量模式
// @author: SaltFish
// @date: 2020/09/08

// Package ch14 is chapter 14
package ch14

import "runtime"

const NCPU = 8

func DoAll() {
	sem := make(chan int, NCPU) // Buffering optional but sensible
	for i := 0; i < NCPU; i++ {
		go DoPart(sem)
	}
	// Drain the channel sem, waiting for NCPU tasks to complete
	for i := 0; i < NCPU; i++ {
		<-sem // wait for one task to complete
	}
	// All done.
}

func DoPart(sem chan int) {
	// do the part of the computation
	sem <- 1 // signal that this piece is done
}

func main() {
	runtime.GOMAXPROCS(NCPU) // runtime.GOMAXPROCS = NCPU
	DoAll()
}
