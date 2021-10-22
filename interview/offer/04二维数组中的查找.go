// @file: 04二维数组中的查找.go
// @date: 2021/2/12

// Package offer
package offer

/*
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
*/

func findNumberIn2DArrayX04(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	x, y := 0, cols-1
	for x < rows && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}
