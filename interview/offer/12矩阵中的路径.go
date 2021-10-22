// @file: 12矩阵中的路径.go
// @date: 2021/2/12

// Package offer
package offer

/*
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

[["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"]]

但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
*/

func existX12(board [][]byte, word string) bool {
	n := len(word)
	rows, cols := len(board), len(board[0])
	visited := make(map[[2]int]bool, 0)
	d := [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}

	var helper func([][]byte, int, int, int) bool
	helper = func(board [][]byte, pos, row, col int) bool {
		if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != word[pos] {
			return false
		}
		if pos == n-1 {
			return true
		}
		for _, v := range d {
			x, y := row+v[0], col+v[1]
			if visited[[2]int{x, y}] {
				continue
			} else {
				visited[[2]int{x, y}] = true
			}
			if helper(board, pos+1, x, y) {
				return true
			}
			visited[[2]int{x, y}] = false
		}
		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			visited[[2]int{i, j}] = true
			if helper(board, 0, i, j) {
				return true
			}
			visited[[2]int{i, j}] = false
		}
	}
	return false
}

func existX12V2(board [][]byte, word string) bool {
	n := len(word)
	rows, cols := len(board), len(board[0])
	d := []int{-1, 0, 1, 0, -1}

	var helper func([][]byte, int, int, int) bool
	helper = func(board [][]byte, pos, row, col int) bool {
		if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != word[pos] {
			return false
		}
		c := board[row][col]
		if c == '#' || c != word[pos] {
			return false
		}
		if pos == n-1 {
			return true
		}
		board[row][col] = '#'
		for i := 0; i < 4; i++ {
			x, y := row+d[i], col+d[i+1]
			if helper(board, pos+1, x, y) {
				return true
			}
		}
		board[row][col] = c
		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if helper(board, 0, i, j) {
				return true
			}
		}
	}
	return false
}
