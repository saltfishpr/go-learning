// @file: dynamic_programming.go
// @date: 2021/1/19

// Package dynamicprogramming
package dynamicprogramming

import "math"

func max(args ...int) int {
	m := math.MinInt64
	for i := range args {
		if args[i] > m {
			m = args[i]
		}
	}
	return m
}

func min(args ...int) int {
	m := math.MaxInt64
	for i := range args {
		if args[i] < m {
			m = args[i]
		}
	}
	return m
}

// 416. 分割等和子集
func canPartition(nums []int) bool {
	return false
}

// 494. 目标和
// sum(A) - sum(B) = target
// sum(A) = target + sum(B)
// sum(A) + sum(A) = target + sum(B) + sum(A)
// 2 * sum(A) = target + sum(nums)
func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum < S || (sum+S)%2 == 1 {
		return 0
	}

	target := (sum + S) / 2
	// dp[i][j] 前i个物品，背包容量为j，有dp[i][j]种方法装满背包
	// 初始化dp[i][0] = 1，dp[0][j] = 0
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
		dp[i][0] = 1
	}

	// dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
	for i := 1; i <= n; i++ {
		for j := 0; j <= target; j++ {
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			} else {
				// 背包空间不足
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][target]
}

func findTargetSumWaysV2(nums []int, S int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum < S || (sum+S)%2 == 1 {
		return 0
	}

	target := (sum + S) / 2
	// 状态压缩
	dp := make([]int, target+1)
	dp[0] = 1

	// dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
	for i := 1; i <= n; i++ {
		// 需要从后向前遍历，保证上一轮i循环dp[j-nums[i-1]]没被覆盖
		for j := target; j >= 0; j-- {
			if j >= nums[i-1] {
				dp[j] = dp[j] + dp[j-nums[i-1]]
			} else {
				// 背包空间不足
				dp[j] = dp[j]
			}
		}
	}
	return dp[target]
}

// 120. 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	if length == 1 {
		return triangle[0][0]
	}
	for i := 1; i < length; i++ {
		triangle[i][0] += triangle[i-1][0]
		for j := 1; j < i; j++ {
			triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
		}
		triangle[i][i] += triangle[i-1][i-1]
	}
	res := math.MaxInt64
	idx := length - 1
	for i := 0; i <= idx; i++ {
		res = min(res, triangle[idx][i])
	}
	return res
}

// 64. 最小路径和
func minPathSum(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	for i := 1; i < rows; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < cols; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			grid[i][j] += min(grid[i][j-1], grid[i-1][j])
		}
	}
	return grid[rows-1][cols-1]
}

// 62. 不同路径
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 63. 不同路径 II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化第一行和第一列
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] != 1 {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
			break
		}
	}
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] != 1 {
			dp[0][i] = 1
		} else {
			dp[0][i] = 0
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			switch {
			case obstacleGrid[i][j] == 1:
				continue
			case obstacleGrid[i-1][j] == 1 && obstacleGrid[i][j-1] == 1:
				dp[i][j] = 0
			case obstacleGrid[i-1][j] == 1:
				dp[i][j] = dp[i][j-1]
			case obstacleGrid[i][j-1] == 1:
				dp[i][j] = dp[i-1][j]
			default:
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// 70. 爬楼梯
func climbStairs(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

// 55. 跳跃游戏
func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if dp[j] == true && j+nums[j] >= i {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}

func canJumpV2(nums []int) bool {
	m := 0
	for i := range nums {
		if m < i {
			return false
		}
		m = max(m, i+nums[i])
	}
	return true
}

// 45. 跳跃游戏 II
func jump(nums []int) int {
	dp := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		dp[i] = math.MaxInt64
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				dp[i] = min(dp[i], dp[j]+1)
			} else {
				continue
			}
		}
	}
	return dp[len(nums)-1]
}

func jumpV2(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

// 5. 最长回文子串
func longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	// 长度
	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = true
			} else if l == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
			if dp[i][j] && l > len(ans)-1 {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}

func longestPalindromeV2(s string) string {
	if s == "" {
		return ""
	}

	expandAroundCenter := func(s string, left, right int) (int, int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left, right = left-1, right+1
		}
		return left + 1, right - 1
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func longestPalindromeV3(s string) string {
	start, end := 0, -1
	t := "#"
	for i := 0; i < len(s); i++ {
		t += string(s[i]) + "#"
	}
	t += "#"
	s = t

	expand := func(s string, left, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left, right = left-1, right+1
		}
		return (right - left - 2) / 2
	}

	armLen := make([]int, 0)
	right, j := -1, -1
	for i := 0; i < len(s); i++ {
		var curArmLen int
		if right >= i {
			iSym := j*2 - i
			minArmLen := min(armLen[iSym], right-i)
			curArmLen = expand(s, i-minArmLen, i+minArmLen)
		} else {
			curArmLen = expand(s, i, i)
		}
		armLen = append(armLen, curArmLen)
		if i+curArmLen > right {
			j = i
			right = i + curArmLen
		}
		if curArmLen*2+1 > end-start {
			start = i - curArmLen
			end = i + curArmLen
		}
	}
	ans := ""
	for i := start; i <= end; i++ {
		if s[i] != '#' {
			ans += string(s[i])
		}
	}
	return ans
}

// 132. 分割回文串 II
func minCut(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			isPalindrome[i][j] = s[i] == s[j] && (j-i <= 1 || isPalindrome[i+1][j-1])
		}
	}

	dp := make([]int, n+1)
	dp[0] = -1
	for i := 1; i <= n; i++ {
		dp[i] = i - 1
		for j := 0; j < i; j++ {
			if isPalindrome[j][i-1] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n]
}

func minCutV2(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i++ {
		for j := i; j < n; j++ {
			isPalindrome[i][j] = s[i] == s[j] && (j-i <= 1 || isPalindrome[i+1][j-1])
		}
	}

	dp := make([]int, n+1)
	dp[0] = -1
	for i := 1; i <= n; i++ {
		dp[i] = i - 1
		for j := 0; j < i; j++ {
			if isPalindrome[j][i-1] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n]
}

func minCutV3(s string) int {
	if len(s) == 0 {
		return 0
	}

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	// 长度
	for i := 0; i < len(s); i++ {
		for start := 0; start < len(s)-i; start++ {
			end := start + i
			if i == 0 {
				dp[start][end] = true
			} else if i == 1 {
				if s[start] == s[end] {
					dp[start][end] = true
				}
			} else {
				if s[start] == s[end] {
					dp[start][end] = dp[start+1][end-1]
				}
			}
		}
	}

	dpIndex := make([]int, len(s)+1)
	for i := range dpIndex {
		dpIndex[i] = len(s)
	}

	dpIndex[0] = 0
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if dp[j][i] {
				if dpIndex[j]+1 < dpIndex[i+1] {
					dpIndex[i+1] = dpIndex[j] + 1
				} else {
					dpIndex[i+1] = dpIndex[i+1]
				}
			}
		}
	}

	return dpIndex[len(s)] - 1
}

// 139. 单词拆分
func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool, 0)
	for i := range wordDict {
		dict[wordDict[i]] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	// 长度
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dict[s[j:i]] && dp[j] == true {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// 1143. 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

// 72. 编辑距离
func minDistance(word1 string, word2 string) int {
	// 初始化dp数组
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	// 边界情况，一个空串和一个非空串的编辑距离为 dp[i][0] = i 和 dp[0][j] = j
	// dp[i][0] 相当于对 word1 执行 i 次删除操作
	for i := 0; i <= len(word1); i++ {
		dp[i][0] = i
	}
	// dp[0][j] 相当于对 word1执行 j 次插入操作。
	for i := 0; i <= len(word2); i++ {
		dp[0][i] = i
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 取插入、删除、替换的最小值
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

// 322. 零钱兑换
func coinChange(coins []int, amount int) int {
	// 金额为i时，组成的最小硬币个数为dp[i]
	// 推导 dp[i]  = min(dp[i-1], dp[i-2], dp[i-5])+1, 前提 i-coins[j] >= 0
	// 初始化为最大值 dp[i]=amount+1
	// 返回值 dp[n] or dp[n]>amount =>-1
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]]+1 < dp[i] {
				dp[i] = dp[i-coins[j]] + 1
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// 92. 背包问题  LintCode
func backPack(m int, A []int) int {
	// dp[i][j] 前i个物品，j容量的背包是否能装下
	// dp[i][j] = dp
	n := len(A)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true

	for i := 1; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			dp[i][j] = dp[i-1][j]
			if j-A[i-1] >= 0 && dp[i-1][j-A[i-1]] {
				dp[i][j] = true
			}
		}
	}

	for i := m; i >= 0; i-- {
		if dp[n][i] {
			return i
		}
	}

	return 0
}

// 125. 背包问题 II LintCode
func backPackII(m int, A []int, V []int) int {
	n := len(A)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			dp[i][j] = dp[i-1][j]
			// 是否加入A[i]物品
			if j-A[i-1] >= 0 && dp[i-1][j-A[i-1]]+V[i-1] > dp[i-1][j] {
				dp[i][j] = dp[i-1][j-A[i-1]] + V[i-1]
			}
		}
	}
	return dp[n][m]
}
