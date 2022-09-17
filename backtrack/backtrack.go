// @file: backtrack.go
// @date: 2021/1/15

// Package backtrack
package backtrack

/*
回溯法
result = []
func backtrack(选择列表,路径):
    if 满足结束条件:
        result.add(路径)
        return
    for 选择 in 选择列表:
        做选择
        backtrack(选择列表,路径)
        撤销选择
*/
// 78. 子集
func subsets(nums []int) [][]int {
	length := len(nums)
	res := make([][]int, 0)

	var backtrack func(int, []int)
	backtrack = func(start int, path []int) {
		// 相当于前序遍历，将选择树的所有节点添加到结果中
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		// 从start开始选择
		for i := start; i < length; i++ {
			path = append(path, nums[i])
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, []int{})
	return res
}

// 77. 组合 [1, n]
func combine(n int, k int) [][]int {
	res := make([][]int, 0)

	var backtrack func(int, []int)
	backtrack = func(start int, path []int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := start; i <= n; i++ {
			path = append(path, i)
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}
	}

	backtrack(1, []int{})
	return res
}

// 46. 全排列
func permute(nums []int) [][]int {
	var res [][]int
	visited := map[int]bool{}

	var backtrack func([]int)
	backtrack = func(path []int) {
		// 结束条件
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for _, num := range nums {
			if visited[num] {
				continue
			}
			// 做选择
			path = append(path, num)
			visited[num] = true
			// 下一个决策
			backtrack(path)
			// 撤销选择
			visited[num] = false
			path = path[:len(path)-1]
		}
	}

	backtrack([]int{})
	return res
}

// 51. N 皇后
func solveNQueens(n int) [][]string {
	var result [][]string
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	isValid := func(board [][]byte, row, col int) bool {
		for i := 0; i < row; i++ {
			if board[i][col] == 'Q' { // 上方
				return false
			}
			if col-row+i >= 0 && board[i][col-row+i] == 'Q' { // 左上方
				return false
			}
			if col+row-i < n && board[i][col+row-i] == 'Q' { // 右上方
				return false
			}
		}
		return true
	}

	var backtrack func([][]byte, int)
	backtrack = func(board [][]byte, row int) {
		if row == len(board) {
			t := make([]string, row)
			for k, bs := range board {
				t[k] = string(bs)
			}
			result = append(result, t)
			return
		}
		for col := 0; col < len(board[row]); col++ {
			if !isValid(board, row, col) {
				continue
			}
			// 做选择
			board[row][col] = 'Q'
			// 进入下一次决策
			backtrack(board, row+1)
			// 撤销选择
			board[row][col] = '.'
		}
	}

	backtrack(board, 0)
	return result
}

// 37. 解数独
func solveSudoku(board [][]byte) {
	isValid := func(board [][]byte, row, col int, c byte) bool {
		x, y := row/3*3, col/3*3
		for i := 0; i < 9; i++ {
			if board[i][col] == c || board[row][i] == c {
				return false
			}
			if board[x+i/3][y+i%3] == c {
				return false
			}
		}
		return true
	}

	var backtrack func([][]byte, int, int) bool
	backtrack = func(board [][]byte, row int, col int) bool {
		// 遍历下一行
		if col == 9 {
			return backtrack(board, row+1, 0)
		}
		// 达到返回情况
		if row == 9 {
			return true
		}
		// 如果该位置是题目给的数字，直接进入下一个位置
		if board[row][col] != '.' {
			return backtrack(board, row, col+1)
		}
		// 对每个数字进行穷举
		for c := byte('1'); c <= '9'; c++ {
			if !isValid(board, row, col, c) {
				continue
			}
			board[row][col] = c
			// 找到可行解立即返回
			if backtrack(board, row, col+1) {
				return true
			}
			// 回溯
			board[row][col] = '.'
		}
		// 穷举完 1~9，依然没有找到可行解，此路不通
		return false

	}

	backtrack(board, 0, 0)
}

// 22. 括号生成
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var backtrack func(int, int, int, []byte)
	backtrack = func(pos, left, right int, bs []byte) {
		if pos == 2*n {
			res = append(res, string(bs))
			return
		}
		if left > 0 {
			bs[pos] = '('
			backtrack(pos+1, left-1, right, bs)
		}
		if left < right {
			bs[pos] = ')'
			backtrack(pos+1, left, right-1, bs)
		}

	}
	backtrack(0, n, n, make([]byte, 2*n))
	return res
}
