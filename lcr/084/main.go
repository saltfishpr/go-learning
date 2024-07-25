package main

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	used := make([]bool, len(nums))

	var res [][]int
	var backtrack func(path []int)
	backtrack = func(path []int) {
		if len(path) == len(nums) {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			res = append(res, pathCopy)
			return
		}

		prevNum := -666
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if prevNum == nums[i] {
				continue
			}
			prevNum = nums[i]

			used[i] = true
			path = append(path, nums[i])
			backtrack(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack(make([]int, 0, len(nums)))
	return res
}
