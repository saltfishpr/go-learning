// @file: 15二进制中1的个数.go
// @date: 2021/2/13

// Package offer
package offer

/*
请实现一个函数，输入一个整数（以二进制串形式），输出该数二进制表示中 1 的个数。例如，把 9表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
*/

func hammingWeightX15(num uint32) int {
	res := 0
	for num != 0 {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}
