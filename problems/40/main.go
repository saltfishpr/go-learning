package main

import "sort"

func main() {
	combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var res [][]int
	var backtrack func(path []int, target int, start int)
	backtrack = func(path []int, target int, start int) {
		if target == 0 {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			res = append(res, pathCopy)
			return
		}
		if target < 0 {
			return
		}
		for i := start; i < len(candidates); i++ {
			// 剪枝，值相同的相邻树枝，只遍历第一条
			if i > start && candidates[i-1] == candidates[i] {
				continue
			}
			path = append(path, candidates[i])
			backtrack(path, target-candidates[i], i+1)
			path = path[:len(path)-1]
		}
	}

	backtrack(make([]int, 0), target, 0)
	return res
}
