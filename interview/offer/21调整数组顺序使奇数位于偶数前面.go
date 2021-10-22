// @file: 21调整数组顺序使奇数位于偶数前面.go
// @date: 2021/2/14

// Package offer
package offer

/*
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
*/

func exchangeX21(nums []int) []int {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		for left < right && nums[left]&1 != 0 {
			left++
		}
		for left < right && nums[right]&1 != 1 {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}
