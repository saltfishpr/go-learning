package main

func combine(n int, k int) [][]int {
	var res [][]int
	var backtrack func(path []int, start int)
	backtrack = func(path []int, start int) {
		if len(path) == k {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			res = append(res, pathCopy)
			return
		}

		for i := start; i < n+1; i++ {
			path = append(path, i)
			backtrack(path, i+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(make([]int, 0, k), 1)
	return res
}
