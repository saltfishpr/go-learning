// @file: stock_trading.go
// @date: 2021/1/17

// Package stocktrading
package stocktrading

import "math"

func initSlice(length int, t interface{}) []interface{} {
	return nil
}

// 121. 买卖股票的最佳时机
func maxProfit1(prices []int) int {
	n := len(prices)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dp0, dp1 := 0, math.MinInt64
	for i := 0; i < n; i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, -prices[i])
	}
	return dp0
}

// 122. 买卖股票的最佳时机 II
func maxProfit2(prices []int) int {
	n := len(prices)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	/*
		dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
		dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
					= max(dp[i-1][k][1], dp[i-1][k][0] - prices[i])
		k = k-1 简化后
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
		可以只保存前一天的值
	*/
	dp0, dp1 := 0, math.MinInt64

	for i := 0; i < n; i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}

	return dp0
}

// 123. 买卖股票的最佳时机 III
func maxProfit3(prices []int) int {
	/*
		dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
		dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])

		dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1]+prices[i])
		dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0]-prices[i])
		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1]+prices[i])
		dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0]-prices[i])
		dp20 = dp[i-1][2][0], dp21 = dp[i-1][2][1], dp10 = dp[i-1][1][0], dp11 = dp[i-1][1][1], dp[i-1][0][0] = 0

		dp20 = max(dp20, dp21+prices[i])
		dp21 = max(dp21, dp10-prices[i])
		dp10 = max(dp10, dp11+prices[i])
		dp11 = max(dp11, -prices[i])
	*/
	n := len(prices)
	dp20, dp21 := 0, math.MinInt64
	dp10, dp11 := 0, math.MinInt64
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 0; i < n; i++ {
		dp20 = max(dp20, dp21+prices[i])
		dp21 = max(dp21, dp10-prices[i])
		dp10 = max(dp10, dp11+prices[i])
		dp11 = max(dp11, -prices[i])
	}
	return dp20
}

// 188. 买卖股票的最佳时机 IV
func maxProfit4(k int, prices []int) int {
	n := len(prices)
	if k >= n>>1 {
		return maxProfit2(prices)
	}
	// 初始化dp数组，[天][交易次数][手上是否持有股票]
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 0; i < n; i++ {
		// base case：
		// dp[-1][k][0] = dp[i][0][0] = 0
		// dp[-1][k][1] = dp[i][0][1] = -infinity
		dp[i][0][0] = 0
		dp[i][0][1] = math.MinInt64
		// 这里不能用外部变量k
		for j := k; j > 0; j-- {
			if i-1 == -1 {
				// max(dp[-1][j][0], dp[-1][j][1]+prices[i]) = max(0, -infinity + prices[i]) = 0
				dp[0][j][0] = 0
				// max(dp[-1][j][1], dp[-1][j-1][0]-prices[i]) = max(-infinity, 0 - prices[i]) = -prices[i]
				dp[0][j][1] = -prices[i]
				continue
			}
			//            max(  选择 rest  ,        选择 sell      )
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			//            max(  选择 rest  ,         选择 buy         )  选择buy时j-1
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}

	return dp[n-1][k][0]
}

// 309. 最佳买卖股票时机含冷冻期
func maxProfit5(prices []int) int {
	n := len(prices)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	dp0, dp1 := 0, math.MinInt64
	dpPre := 0 // dp[i-2][0]

	for i := 0; i < n; i++ {
		dp0, dp1, dpPre = max(dp0, dp1+prices[i]), max(dp1, dpPre-prices[i]), dp0
	}

	return dp0
}

// 714. 买卖股票的最佳时机含手续费
func maxProfit6(prices []int, fee int) int {
	n := len(prices)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	dp0, dp1 := 0, math.MinInt64

	for i := 0; i < n; i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i]-fee)
	}

	return dp0
}
