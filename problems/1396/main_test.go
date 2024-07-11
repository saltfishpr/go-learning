package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUndergroundSystem(t *testing.T) {
	t.Run("示例 1", func(t *testing.T) {
		us := Constructor()

		us.CheckIn(45, "Leyton", 3)
		us.CheckIn(32, "Paradise", 8)
		us.CheckIn(27, "Leyton", 10)
		us.CheckOut(45, "Waterloo", 15)
		us.CheckOut(27, "Waterloo", 20)
		us.CheckOut(32, "Cambridge", 22)
		assert.Equal(t, 14., us.GetAverageTime("Paradise", "Cambridge"))
		assert.Equal(t, 11., us.GetAverageTime("Leyton", "Waterloo"))
		us.CheckIn(10, "Leyton", 24)
		assert.Equal(t, 11., us.GetAverageTime("Leyton", "Waterloo"))
		us.CheckOut(10, "Waterloo", 38)
		assert.Equal(t, 12., us.GetAverageTime("Leyton", "Waterloo"))
	})

	t.Run("示例 2", func(t *testing.T) {
		us := Constructor()

		us.CheckIn(10, "Leyton", 3)
		us.CheckOut(10, "Paradise", 8)
		assert.Equal(t, 5./1, us.GetAverageTime("Leyton", "Paradise"))
		us.CheckIn(5, "Leyton", 10)
		us.CheckOut(5, "Paradise", 16)
		assert.Equal(t, 11./2, us.GetAverageTime("Leyton", "Paradise"))

		us.CheckIn(2, "Leyton", 21)
		us.CheckOut(2, "Paradise", 30)
		assert.Equal(t, 20./3, us.GetAverageTime("Leyton", "Paradise"))
	})
}

var result float64

func BenchmarkUndergroundSystem(b *testing.B) {
	var res float64
	us := Constructor()
	startStation := "A"
	endStation := "B"
	for n := 0; n < b.N; n++ {
		us.CheckIn(n, startStation, 5)
		res = us.GetAverageTime(startStation, endStation)
		us.CheckOut(n, endStation, 10)
	}
	result = res
}
