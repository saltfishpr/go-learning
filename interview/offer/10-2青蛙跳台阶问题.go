// @file: 10-2青蛙跳台阶问题.go
// @date: 2021/2/12

// Package offer
package offer

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
*/

func numWaysX10(n int) int {
	a, b := 0, 1
	for n > 0 {
		a, b = b, (a+b)%(1e9+7)
		n--
	}
	return b
}
