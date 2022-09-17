// @file: 53-2缺失的数字.go
// @date: 2021/2/23

// Package offer
package offer

/*
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
*/

func missingNumberX53(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
