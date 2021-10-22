// @file: 17打印从1到最大的n位数.go
// @date: 2021/2/13

// Package offer
package offer

/*
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
*/

func printNumbersX17(n int) []int {
	cnt := 1
	for i := 0; i < n; i++ {
		cnt *= 10
	}
	res := make([]int, cnt-1)
	for i := 0; i < cnt-1; i++ {
		res[i] = i + 1
	}
	return res
}
