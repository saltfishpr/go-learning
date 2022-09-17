// @file: rob_house.go
// @date: 2021/1/17

// Package robhouse
package robhouse

// 198. 打家劫舍
func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}

// 213. 打家劫舍 II
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	return max(rob1(nums[:n-1]), rob1(nums[1:]))
}

// 337. 打家劫舍 III
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob3(root *TreeNode) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	var dp func(*TreeNode) (int, int)
	dp = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		l1, l2 := dp(root.Left)
		r1, r2 := dp(root.Right)

		不干 := max(l1, l2) + max(r1, r2)
		干 := root.Val + l1 + r1
		return 不干, 干
	}

	return max(dp(root))
}
