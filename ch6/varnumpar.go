// @file: varnumpar.go
// @description: 变长参数
// @author: SaltFish
// @date: 2020/08/04

// Package ch6 is chapter 6
package ch6

// MyMin 返回int slice中的最小值
func MyMin(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
