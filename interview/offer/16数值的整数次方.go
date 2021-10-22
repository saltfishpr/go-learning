// @file: 16数值的整数次方.go
// @date: 2021/2/13

// Package offer
package offer

/*
实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。
*/

func myPowX16(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n < 0 {
		x, n = 1/x, -n
	}
	var res float64 = 1
	for n != 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}
