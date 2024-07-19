package main

func permute(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))

	var backtrack func(nums []int, path []int)
	backtrack = func(nums, path []int) {
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

			path = append(path, num)
			used[i] = true
			backtrack(nums, path)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	backtrack(nums, make([]int, 0, len(nums)))
	return res
}
