package main

func permute(nums []int) [][]int {
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

		for i, num := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, num)
			backtrack(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack(make([]int, 0, len(nums)))
	return res
}
