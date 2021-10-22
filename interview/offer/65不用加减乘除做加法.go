// @file: 65不用加减乘除做加法.go
// @date: 2021/2/26
// Package offer

package offer

/*
写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。
*/

func addX65(a int, b int) int {
	for b != 0 {
		c := (a & b) << 1 // 进位
		a ^= b            // A保存和
		b = c             // B保存每一位的进位
		// 循环直到没有进位
	}
	return a
}
