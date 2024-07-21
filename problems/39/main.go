package main

func combinationSum(candidates []int, target int) [][]int {
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
			path = append(path, candidates[i])
			backtrack(path, target-candidates[i], i)
			path = path[:len(path)-1]
		}
	}

	backtrack(make([]int, 0), target, 0)
	return res
}
