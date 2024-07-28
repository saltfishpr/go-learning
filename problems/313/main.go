package main

func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	idx := make([]int, len(primes))
	for i := 0; i < len(primes); i++ {
		idx[i] = 1
	}
	nums := make([]int, len(primes))

	for i := 2; i < n+1; i++ {
		for j := 0; j < len(primes); j++ {
			nums[j] = dp[idx[j]] * primes[j]
		}
		dp[i] = min(nums...)
		for j := 0; j < len(primes); j++ {
			if nums[j] == dp[i] {
				idx[j]++
			}
		}
	}
	return dp[n]
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	res := s[0]
	for _, v := range s {
		if v < res {
			res = v
		}
	}
	return res
}
