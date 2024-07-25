package main

func combinationSum(candidates []int, target int) [][]int {
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
			path = append(path, candidates[i])
			backtrack(path, i, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	backtrack(make([]int, 0), 0, 0)
	return res
}
