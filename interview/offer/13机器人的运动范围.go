// @file: 13机器人的运动范围.go
// @date: 2021/2/13

// Package offer
package offer

/*
地上有一个m行n列的方格，从坐标 [0, 0] 到坐标 [m-1, n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
*/

func movingCountX13(m int, n int, k int) int {
	res := 0
	d := []int{-1, 0, 1, 0, -1}
	board := make([][]bool, m)

	for i := 0; i < m; i++ {
		board[i] = make([]bool, n)
	}

	check := func(row, col int) bool {
		res := 0
		for row != 0 {
			res += row % 10
			row /= 10
		}
		for col != 0 {
			res += col % 10
			col /= 10
		}
		return res <= k
	}

	var helper func([][]bool, int, int)
	helper = func(board [][]bool, row int, col int) {
		if row < 0 || row >= m || col < 0 || col >= n {
			return
		}
		if board[row][col] || !check(row, col) {
			return
		}
		board[row][col] = true
		res++
		for i := 0; i < 4; i++ {
			x, y := row+d[i], col+d[i+1]
			helper(board, x, y)
		}
	}

	helper(board, 0, 0)
	return res
}
