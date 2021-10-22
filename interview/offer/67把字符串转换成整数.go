// @file: 67把字符串转换成整数.go
// @date: 2021/2/26

// Package offer
package offer

import (
	"math"
	"strings"
)

/*
写一个函数 StrToInt，实现把字符串转换成整数这个功能。不能使用 atoi 或者其他类似的库函数。
*/

func strToIntX67(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	res, i, sign := 0, 1, true
	intMax, intMin, bndry := math.MaxInt32, math.MinInt32, math.MaxInt32/10
	if str[0] == '-' {
		sign = false
	} else if str[0] != '+' {
		i = 0
	}
	str = str[i:]
	for k := 0; k < len(str); k++ {
		c := str[k]
		if c < '0' || c > '9' {
			break
		}
		if res > bndry || res == bndry && c > '7' {
			if sign {
				return intMax
			} else {
				return intMin
			}
		}
		res = 10*res + int(c-'0')
	}
	if sign {
		return res
	} else {
		return -res
	}
}
