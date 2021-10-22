// Package offer
// @file: 42连续子数组的最大和.go
// @date: 2021/2/18
package offer

/*
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值
要求时间复杂度为O(n)。
*/

func maxSubArrayX42(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] += max(0, nums[i-1])
		res = max(res, nums[i])
	}
	return res
}
