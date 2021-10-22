// @file: 39数组中出现次数超过一半的数字.go
// @date: 2021/2/18

// Package offer
package offer

/*
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/

func majorityElementX39(nums []int) int {
	votes, res := 0, 0
	for i := 0; i < len(nums); i++ {
		if votes == 0 {
			res = nums[i]
		}
		if res == nums[i] {
			votes++
		} else {
			votes--
		}
	}
	return res
}
