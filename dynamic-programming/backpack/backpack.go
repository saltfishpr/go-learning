// @file: backpack.go
// @date: 2021/3/7

// Package backpack
package backpack

// 416. 分割等和子集
func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2

	// dp[i][j] 前i个数，背包容量为j
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i] >= 0 {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else {
				// 容量不足，不能装入第i个物品
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][sum]
}

func canPartitionV2(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2

	// dp[i][j] 前i个数，背包容量为j
	dp := make([]bool, sum+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := sum; j >= 0; j-- {
			if j-nums[i-1] >= 0 {
				dp[j] = dp[j] || dp[j-nums[i-1]]
			}
		}
	}
	return dp[sum]
}

// 518. 零钱兑换 II
func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i] < 0 {
				continue
			}
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}
