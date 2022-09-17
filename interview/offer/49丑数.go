// @file: 49丑数.go
// @date: 2021/2/21

// Package offer
package offer

/*
我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。
*/

func nthUglyNumberX49(n int) int {
	dp := make([]int, n)

	min := func(args ...int) int {
		minVal := args[0]
		for i := 1; i < len(args); i++ {
			if args[i] < minVal {
				minVal = args[i]
			}
		}
		return minVal
	}

	dp[0] = 1
	a, b, c := 0, 0, 0

	for i := 1; i < n; i++ {
		n2, n3, n5 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = min(n2, n3, n5)
		if dp[i] == n2 {
			a++
		}
		if dp[i] == n3 {
			b++
		}
		if dp[i] == n5 {
			c++
		}
	}
	return dp[n-1]
}
