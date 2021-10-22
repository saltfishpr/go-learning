// @file: 63股票的最大利润.go
// @date: 2021/2/26
// Package offer

package offer

/*
假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
*/

func maxProfitX63(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	dp := make([]int, n)
	cost := prices[0]
	for i := 1; i < n; i++ {
		cost = min(cost, prices[i])
		dp[i] = max(dp[i-1], prices[i]-cost)
	}
	return dp[n-1]
}

func maxProfitX63V2(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	profit := 0
	cost := prices[0]
	for i := 0; i < n; i++ {
		cost = min(cost, prices[i])
		profit = max(profit, prices[i]-cost)
	}
	return profit
}
