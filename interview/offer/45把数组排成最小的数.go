// @file: 45把数组排成最小的数.go
// @date: 2021/2/19

// Package offer
package offer

import (
	"strconv"
)

/*
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
*/

func minNumberX45(nums []int) string {
	var fastSort func(int, int)
	fastSort = func(l int, r int) {
		if l >= r {
			return
		}
		i, j := l, r
		for i < j {
			for i < j && strconv.Itoa(nums[j])+strconv.Itoa(nums[l]) >= strconv.Itoa(nums[l])+strconv.Itoa(nums[j]) {
				j--
			}
			for i < j && strconv.Itoa(nums[i])+strconv.Itoa(nums[l]) <= strconv.Itoa(nums[l])+strconv.Itoa(nums[i]) {
				i++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[i], nums[l] = nums[l], nums[i]
		fastSort(l, i-1)
		fastSort(i+1, r)
	}

	fastSort(0, len(nums)-1)
	res := ""
	for i := 0; i < len(nums); i++ {
		res += strconv.Itoa(nums[i])
	}
	return res
}
