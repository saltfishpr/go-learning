// @file: 03数组中重复的数字.go
// @date: 2021/2/12

// Package offer
package offer

/*
找出数组中重复的数字。

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
*/

func findRepeatNumberX03(nums []int) int {
	n := len(nums)
	counter := make([]int, n)

	for _, v := range nums {
		if counter[v] != 0 {
			return v
		}
		counter[v]++
	}
	return 0
}

func findRepeatNumberX03V2(nums []int) int {
	counter := make(map[int]struct{}, 0)

	for _, v := range nums {
		if _, ok := counter[v]; ok {
			return v
		}
		counter[v] = struct{}{}
	}
	return 0
}
