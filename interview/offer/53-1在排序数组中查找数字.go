// @file: 53-1在排序数组中查找数字.go
// @date: 2021/2/23

// Package offer
package offer

func searchX53(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 先查找左边界
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			right = mid - 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return 0
	}
	res := 0
	for i := left; i < len(nums); i++ {
		if nums[i] == target {
			res++
		} else {
			break
		}
	}
	return res
}
