// @file: others.go
// @date: 2021/3/8

// Package others
package others

import (
	"math"
)

// 10. 正则表达式匹配
func isMatch(s string, p string) bool {
	n1, n2 := len(s), len(p)
	dp := make([][]bool, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]bool, n2+1)
	}

	dp[0][0] = true
	for i := 2; i <= n2; i += 2 {
		// *前面的字符全部为0次
		dp[0][i] = dp[0][i-2] && p[i-1] == '*'
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || dp[i][j-1] || dp[i-1][j] && s[i-1] == p[j-2] ||
					dp[i-1][j] && p[j-2] == '.'
			} else {
				dp[i][j] = dp[i-1][j-1] && s[i-1] == p[j-1] || dp[i-1][j-1] && p[j-1] == '.'
			}
		}
	}

	return dp[n1][n2]
}

// 887. 鸡蛋掉落
func superEggDrop(k int, n int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// dp[k][n] k个鸡蛋，n层楼，最少要试多少次
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n+1)
	}
	// 初始情况，没有鸡蛋则试0次，楼层为0则试0次
	// 只有一个鸡蛋，则只能遍历，尝试次数为楼层数
	for i := 0; i <= n; i++ {
		dp[1][i] = i
	}

	// 从两个鸡蛋开始填充dp数组
	for i := 2; i <= k; i++ {
		for j := 1; j <= n; j++ {
			/*
				y := math.MaxInt64
				// 从第一层试到第j层
				for x := 1; x <= j; x++ {
					y = min(y, max(dp[i][j-x], dp[i-1][x-1])+1)
				}
				dp[i][j] = y
			*/

			// 二分查找优化
			left, right := 1, j
			res := math.MaxInt64
			for left <= right {
				x := (left + right) / 2
				// dp[k-1][x-1] 是 x 的单调递增函数
				// dp[k][j-x] 是 x 的单调递减函数
				// 所以可以使用二分查找优化，找到两个函数的交点左右两个值为 left, right
				t1, t2 := dp[i-1][x-1], dp[i][j-x]
				// res = min(max(碎，没碎) + 1)
				if t1 < t2 {
					left = x + 1
					res = min(res, t2+1)
				} else {
					right = x - 1
					res = min(res, t1+1)
				}
			}
			dp[i][j] = res

		}
	}
	return dp[k][n]
}

func superEggDropV2(K int, N int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = i
	}
	dp2 := make([]int, N+1)

	for k := 2; k < K+1; k++ {
		x := 1
		for n := 1; n < N+1; n++ {
			for x < n && max(dp[x-1], dp2[n-x]) >= max(dp[x], dp2[n-x-1]) {
				x++
			}
			dp2[n] = 1 + max(dp[x-1], dp2[n-x])
		}
		copy(dp, dp2)
	}
	return dp[N]
}

func superEggDropV3(K int, N int) int {
	if N == 1 {
		return 1
	}
	// dp[i][j] 可以做 i 次操作，有 j 个鸡蛋
	dp := make([][]int, N+1) // 最多可以做 N 次操作
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, K+1)
	}
	for i := 1; i < K+1; i++ {
		dp[1][i] = 1
	}
	res := -1
	for i := 2; i < N+1; i++ {
		for j := 1; j < K+1; j++ {
			dp[i][j] = 1 + dp[i-1][j-1] + dp[i-1][j]
		}
		if dp[i][K] >= N {
			res = i
			break
		}
	}
	return res
}

// 312. 戳气球
func maxCoins(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(nums)
	points := make([]int, n+2)
	points[0], points[n+1] = 1, 1
	copy(points[1:n+1], nums)

	// dp[i][j]表示戳破气球i和气球j之间（开区间，不包括i和j）的所有气球，可以获得的最高分数为x
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}

	for i := n; i >= 0; i-- {
		for j := i + 1; j < n+2; j++ {
			// i, j 为左右边界气球🎈，可取到0和n+1
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+points[k]*points[i]*points[j])
			}
		}
	}
	return dp[0][n+1]
}

// 877. 石子游戏
func stoneGame(piles []int) bool {
	n := len(piles)
	// dp[i][j][0] 表示[i, j]石头的子序列 先手 能获得的最高分数
	// dp[i][j][1] 表示[i, j]石头的子序列 后手 能获得的最高分数
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, n)
	}

	// base case
	for i := 0; i < n; i++ {
		dp[i][i][0] = piles[i]
		dp[i][i][1] = 0
	}

	// l长度，斜着遍历数组
	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := l + i - 1
			left := piles[i] + dp[i+1][j][1]  // 先手 选择左边的分数
			right := piles[j] + dp[i][j-1][1] // 先手 选择右边的分数
			if left > right {
				dp[i][j][0] = left
				dp[i][j][1] = dp[i+1][j][0]
			} else {
				dp[i][j][0] = right
				dp[i][j][1] = dp[i][j-1][0]
			}
		}
	}
	return dp[0][n-1][0] > dp[0][n-1][1]
}

// 651. 四键键盘
func maxA(n int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		for j := 2; j < i; j++ {
			dp[i] = max(dp[i], dp[i-2]*(i-j+1))
		}
	}

	return dp[n]
}

// time: O(m+n), space: O(m)
func kmp(pat, txt string) int {
	m := len(pat)
	n := len(txt)
	// dp[状态][字符] = 下个状态
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, 256)
	}
	// base case，在状态0只有匹配到pat[0]才能进入下一个状态
	dp[0][pat[0]] = 1
	// 构建dp数组
	x := 0 // 影子状态初始为0
	for i := 1; i < m; i++ {
		for c := 0; c < 256; c++ {
			dp[i][c] = dp[x][c]
		}
		dp[i][pat[i]] = i + 1
		x = dp[x][pat[i]]
	}
	// 根据dp数组在txt中查找
	j := 0
	for i := 0; i < n; i++ {
		j = dp[j][txt[i]]
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

// 1312. 让字符串成为回文串的最少插入次数
func minInsertions(s string) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][n-1]
}

func minInsertionsV2(s string) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(s)
	dp := make([]int, n)

	temp := 0
	for i := n - 2; i >= 0; i-- {
		pre := 0
		for j := i + 1; j < n; j++ {
			temp = dp[j]
			if s[i] == s[j] {
				dp[j] = pre
			} else {
				dp[j] = min(dp[j], dp[j-1]) + 1
			}
			pre = temp
		}
	}
	return dp[n-1]
}
