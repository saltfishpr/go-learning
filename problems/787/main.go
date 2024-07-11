package main

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 10000*101 + 1

	// dp[t][i] 表示通过恰好 t 次航班，从出发城市 src 到达城市 i 需要的最小花费
	// k 次中转，也就是最多搭乘 k+1 次航班，再加上 base state，0 <= t < k+2
	dp := make([][]int, k+2)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[0][src] = 0

	for t := 1; t < k+2; t++ {
		for _, flight := range flights {
			j, i, cost := flight[0], flight[1], flight[2]
			dp[t][i] = min(dp[t][i], dp[t-1][j]+cost)
		}
	}

	ans := inf
	for t := 1; t <= k+1; t++ {
		ans = min(ans, dp[t][dst])
	}
	if ans == inf {
		ans = -1
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
