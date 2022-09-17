// @file: 47礼物的最大价值.go
// @date: 2021/2/21

// Package offer
package offer

/*
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
*/

func maxValueX47(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dp := make([][]int, rows)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < cols; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[rows-1][cols-1]
}
