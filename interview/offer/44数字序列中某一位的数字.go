// @file: 44数字序列中某一位的数字.go
// @date: 2021/2/19

// Package offer
package offer

import "strconv"

/*
数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。

请写一个函数，求任意第n位对应的数字。
*/

func findNthDigitX44(n int) int {
	digit, start, count := 1, 1, 9
	// 计算是几位数字
	for count < n {
		n -= count
		start *= 10
		digit += 1
		count = 9 * start * digit
	}
	// n位所在的数
	num := start + (n-1)/digit
	s := strconv.Itoa(num)
	// n位对应的数字
	res := int(s[(n-1)%digit] - '0')
	return res
}
