package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var res [][]int
	var backtrack func(path []int, start int, sum int)
	backtrack = func(path []int, start int, sum int) {
		if sum == target {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			res = append(res, pathCopy)
			return
		}
		if sum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			backtrack(path, i+1, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(make([]int, 0), 0, 0)
	return res
}
