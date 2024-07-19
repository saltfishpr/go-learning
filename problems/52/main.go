package main

func main() {
	println(totalNQueens(16))
}

func totalNQueens(n int) int {
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	isValid := func(i, j int) bool {
		// 上方
		for k := 0; k < i; k++ {
			if board[k][j] == 'Q' {
				return false
			}
		}
		// 右上
		for k1, k2 := i, j; k1 >= 0 && k2 < n; k1, k2 = k1-1, k2+1 {
			if board[k1][k2] == 'Q' {
				return false
			}
		}
		// 左上
		for k1, k2 := i, j; k1 >= 0 && k2 >= 0; k1, k2 = k1-1, k2-1 {
			if board[k1][k2] == 'Q' {
				return false
			}
		}
		return true
	}

	var res int
	var backtrack func(i int)
	backtrack = func(i int) {
		if i == n {
			res++
			return
		}

		for j := 0; j < n; j++ {
			if !isValid(i, j) {
				continue
			}
			board[i][j] = 'Q'
			backtrack(i + 1)
			board[i][j] = '.'
		}
	}

	backtrack(0)
	return res
}
