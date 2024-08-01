package main

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i < n+1; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], j*(i-j), j*dp[i-j])
		}
	}
	return dp[n]
}
