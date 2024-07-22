package main

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var backtrack func(path []int, start int, sum int)
	backtrack = func(path []int, start, sum int) {
		if len(path) == k {
			if sum == n {
				pathCopy := make([]int, len(path))
				copy(pathCopy, path)
				res = append(res, pathCopy)
			}
			return
		}

		for i := start; i < 10; i++ {
			path = append(path, i)
			backtrack(path, i+1, sum+i)
			path = path[:len(path)-1]
		}
	}

	backtrack(make([]int, 0, k), 1, 0)
	return res
}
