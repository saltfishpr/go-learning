// @file: 57-1和为s的两个数字.go
// @date: 2021/2/23

// Package offer
package offer

/*
输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。
*/

func twoSumX57(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		left, right := nums[i], nums[j]
		sum := left + right
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			for i < j && nums[i] == left {
				i++
			}
		} else {
			for i < j && nums[j] == right {
				j--
			}
		}
	}
	return []int{}
}
