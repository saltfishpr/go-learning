package main

func numDistinct(s string, t string) int {
	m, n := len(t), len(s)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = 1
	}
	// t[:i] 和 s[:j]
	for i := 1; i < m+1; i++ {
		for j := i; j < n+1; j++ {
			if s[j-1] == t[i-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}
