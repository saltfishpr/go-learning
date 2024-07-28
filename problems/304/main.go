package main

type NumMatrix struct {
	data [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	data := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		data[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(data[i]); j++ {
			data[i][j] = sum(data, i-1, j) + sum(data, i, j-1) - sum(data, i-1, j-1) + matrix[i][j]
		}
	}
	return NumMatrix{
		data: data,
	}
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return sum(nm.data, row2, col2) - sum(nm.data, row1-1, col2) - sum(nm.data, row2, col1-1) + sum(nm.data, row1-1, col1-1)
}

func sum(matrix [][]int, x, y int) int {
	if x < 0 || y < 0 {
		return 0
	}
	return matrix[x][y]
}
