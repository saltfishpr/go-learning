package main

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i < n+1; i++ {
		num2 := dp[p2] * 2
		num3 := dp[p3] * 3
		num5 := dp[p5] * 5

		dp[i] = min(num2, num3, num5)
		if dp[i] == num2 {
			p2++
		}
		if dp[i] == num3 {
			p3++
		}
		if dp[i] == num5 {
			p5++
		}
	}
	return dp[n]
}
