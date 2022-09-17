// @file: 46把数字翻译成字符串.go
// @date: 2021/2/19

// Package offer
package offer

import "strconv"

/*
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。
请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
*/

func translateNumX46(num int) int {
	str := strconv.Itoa(num)
	n := len(str)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 1; i < n; i++ {
		pre, cur := int(str[i-1]-'0'), int(str[i]-'0')
		if pre != 0 && pre*10+cur < 26 {
			dp[i+1] = dp[i-1] + dp[i]
		} else {
			dp[i+1] = dp[i]
		}
	}
	return dp[n]
}
