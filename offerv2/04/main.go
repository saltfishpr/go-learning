package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	i, j := 0, n-1
	for i <= m-1 && j >= 0 {
		// 从右上角开始搜索
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
			continue
		} else if matrix[i][j] < target {
			i++
			continue
		}
	}

	return false
}
