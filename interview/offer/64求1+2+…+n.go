// @file: 64求1+2+…+n.go
// @date: 2021/2/26

// Package offer
package offer

/*
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
*/

func sumNumsX64(n int) int {
	res := 0
	var helper func(int) bool
	helper = func(n int) bool {
		_ = n > 1 && helper(n-1)
		res += n
		return true
	}

	helper(n)
	return res
}
