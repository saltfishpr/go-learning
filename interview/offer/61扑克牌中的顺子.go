// @file: 61扑克牌中的顺子.go
// @date: 2021/2/24

// Package offer
package offer

import "sort"

/*
从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。
2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。
*/

func isStraightX61(nums []int) bool {
	sort.Ints(nums)
	joker := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			joker++
		} else {
			break
		}
	}
	if nums[len(nums)-1] >= nums[joker]+5 {
		return false
	}

	for i := joker; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return false
		}
	}
	return true
}
