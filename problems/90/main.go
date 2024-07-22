package main

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	var backtrack func(path []int, start int)
	backtrack = func(path []int, start int) {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		res = append(res, pathCopy)

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtrack(path, i+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(make([]int, 0, len(nums)), 0)
	return res
}
