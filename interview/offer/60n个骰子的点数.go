// @file: 60n个骰子的点数.go
// @date: 2021/2/24

// Package offer
package offer

import "math"

/*
把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。

你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。
*/

func dicesProbabilityX60(n int) []float64 {
	dp := make([]int, 70)
	for i := 1; i <= 6; i++ {
		dp[i] = 1
	}
	// j 个骰子
	for j := 2; j <= n; j++ {
		// 点数
		for i := 6 * j; i >= j; i-- {
			dp[i] = 0
			for cur := 1; cur <= 6; cur++ {
				if i-cur < j-1 {
					break
				}
				dp[i] += dp[i-cur]
			}
		}
	}

	all := math.Pow(6, float64(n))
	res := make([]float64, 0)
	for i := n; i <= 6*n; i++ {
		res = append(res, float64(dp[i])/all)
	}

	return res
}
