package main

func solveNQueens(n int) [][]string {
	// 初始化棋盘
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	// 检查 i, j 上放置皇后是否有冲突
	isValid := func(i, j int) bool {
		// 列是否有冲突
		for k := 0; k < i; k++ {
			if board[k][j] == 'Q' {
				return false
			}
		}
		// 右上方是否有冲突
		for k1, k2 := i-1, j+1; k1 >= 0 && k2 < n; k1, k2 = k1-1, k2+1 {
			if board[k1][k2] == 'Q' {
				return false
			}
		}
		// 左上方是否有冲突
		for k1, k2 := i-1, j-1; k1 >= 0 && k2 >= 0; k1, k2 = k1-1, k2-1 {
			if board[k1][k2] == 'Q' {
				return false
			}
		}
		return true
	}

	var res [][]string
	var backtrack func(path []string, i int)
	backtrack = func(path []string, i int) {
		if i == n {
			pathCopy := make([]string, len(path))
			copy(pathCopy, path)
			res = append(res, pathCopy)
			return
		}
		for j := 0; j < n; j++ {
			if !isValid(i, j) {
				continue
			}
			board[i][j] = 'Q'
			backtrack(append(path, string(board[i])), i+1)
			board[i][j] = '.'
		}
	}

	backtrack(make([]string, 0, n), 0)
	return res
}
