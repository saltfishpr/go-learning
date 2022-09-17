// @file: 59-2滑动窗口的最大值.go
// @date: 2021/2/24

// Package offer
package offer

/*
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。
*/

func maxSlidingWindowX59(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	queue := make([]int, 0)
	res := make([]int, len(nums)-k+1)

	for i := 0; i < k; i++ {
		for len(queue) != 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
	}
	res[0] = queue[0]
	for i := k; i < len(nums); i++ {
		// 滑动窗口左侧的值为单调队列最大值，则最大值出队列
		if queue[0] == nums[i-k] {
			queue = queue[1:]
		}
		// 新进入滑动窗口的值为 nums[i]
		// 从后面弹出队列中小于 nums[i] 的所有值
		for len(queue) != 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		res[i-k+1] = queue[0]
	}
	return res
}
