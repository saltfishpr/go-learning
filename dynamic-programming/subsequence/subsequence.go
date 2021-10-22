// @file: subsequence.go
// @date: 2021/3/5

// Package subsequence
package subsequence

import (
	"sort"
)

// 72. 编辑距离
func minDistance(word1 string, word2 string) int {
	min := func(a int, others ...int) int {
		minVal := a
		for i := 0; i < len(others); i++ {
			if others[i] < minVal {
				minVal = others[i]
			}
		}
		return minVal
	}

	// dp[i][j] word1[:i]和word2[:j]的编辑距离
	n1, n2 := len(word1), len(word2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}

	for i := 0; i <= n1; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= n2; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				// 字母相同就什么都不做
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j]+1, dp[i-1][j-1]+1) // 插入 删除 替换
			}
		}
	}

	return dp[n1][n2]
}

// 354. 俄罗斯套娃信封问题
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	sort.Slice(
		envelopes, func(i, j int) bool {
			a, b := envelopes[i], envelopes[j]
			// 宽度递增，高度递减排序
			return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
		},
	)

	// dp[i]表示第i个信封里最多可以装几个信封
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}

func maxEnvelopesV2(envelopes [][]int) int {
	sort.Slice(
		envelopes, func(i, j int) bool {
			a, b := envelopes[i], envelopes[j]
			return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
		},
	)

	// 对 height 寻找最长递增子序列
	dp := make([]int, 0)
	for _, envelope := range envelopes {
		h := envelope[1]
		if i := sort.SearchInts(dp, h); i < len(dp) {
			dp[i] = h
		} else {
			dp = append(dp, h)
		}
	}
	// 返回最长递增子序列长度
	return len(dp)
}

// 53. 最大子序和
func maxSubArray(nums []int) int {
	n := len(nums)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// dp[i] 为以 nums[i] 为结尾的最大子数组和
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		// 自成一派 / 跟前面结合
		dp[i] = max(nums[i], nums[i]+dp[i-1])
	}

	res := dp[0]
	for i := 1; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}

func maxSubArrayV2(nums []int) int {
	n := len(nums)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// dp[i] 为以 nums[i] 为结尾的最大子数组和
	dp0, dp1, res := nums[0], 0, nums[0]
	for i := 1; i < n; i++ {
		// 自成一派 / 跟前面结合
		dp1 = max(nums[i], nums[i]+dp0)
		dp0 = dp1
		res = max(res, dp1)
	}

	return res
}

// 1143. 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n1, n2 := len(text1), len(text2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n1][n2]
}

// 583. 两个字符串的删除操作
func minDistanceX583(word1 string, word2 string) int {
	// 删除后的答案就是最长公共子序列
	lcs := longestCommonSubsequence(word1, word2)
	return len(word1) + len(word2) - lcs - lcs
}

// 712. 两个字符串的最小ASCII删除和
func minimumDeleteSum(s1 string, s2 string) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n1, n2 := len(s1), len(s2)
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for i := 1; i <= n2; i++ {
		dp[0][i] = dp[0][i-1] + int(s2[i-1])
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}

	return dp[n1][n2]
}

// 516. 最长回文子序列
func longestPalindromeSubseq(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(s)
	// dp[i][j] 表示 s[i:j]的最长回文子序列长度
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	// 初始情况
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	// dp[i][j] 和 dp[i+1][j-1]有关
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
