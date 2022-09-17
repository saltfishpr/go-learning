// @file: bfs.go
// @date: 2021/1/15

// Package bfs
package bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 111. 二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	queue = append(queue, root)
	depth := 1
	for len(queue) != 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}

// 752. 打开转盘锁
func openLock(deadends []string, target string) int {
	type lock [4]byte

	var queue []lock
	var start = lock{'0', '0', '0', '0'}
	var visited = map[lock]bool{}

	// 字符串转lock
	stringToLock := func(s string) lock {
		var res lock
		for i := range s {
			res[i] = s[i]
		}
		return res
	}
	// 判断两个lock类型是否相等
	isLockEqual := func(l1, l2 lock) bool {
		for i := range l1 {
			if l1[i] != l2[i] {
				return false
			}
		}
		return true
	}

	var targetLock = stringToLock(target)

	// 转动增加一位数字
	plusOne := func(l lock, i int) lock {
		if l[i] == '9' {
			l[i] = '0'
			return l
		}
		l[i]++
		return l
	}
	// 转动减少一位数字
	minusOne := func(l lock, i int) lock {
		if l[i] == '0' {
			l[i] = '9'
			return l
		}
		l[i]--
		return l
	}

	for _, v := range deadends {
		visited[stringToLock(v)] = true
	}
	if visited[start] {
		return -1
	}

	// 初始化
	queue = append(queue, start)
	visited[start] = true
	step := 0

	for len(queue) != 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			lock1 := queue[0]
			queue = queue[1:]

			if isLockEqual(lock1, targetLock) {
				return step
			}

			for j := 0; j < 4; j++ {
				// 拷贝
				up := plusOne(lock1, j)
				if visited[up] != true {
					queue = append(queue, up)
					visited[up] = true
				}
				down := minusOne(lock1, j)
				if visited[down] != true {
					queue = append(queue, down)
					visited[down] = true
				}
			}
		}
		step++
	}
	return -1
}

// 773. 滑动谜题
func slidingPuzzle(board [][]int) int {
	neighbor := [][]int{
		{1, 3},
		{0, 4, 2},
		{1, 5},
		{0, 4},
		{3, 1, 5},
		{4, 2},
	}

	start := make([]byte, 6)
	for i := range board {
		for j := range board[i] {
			start[i*3+j] = byte(board[i][j] + '0')
		}
	}
	target := "123450"
	visited := make(map[string]struct{}, 0)
	queue := make([][]byte, 0)
	step := 0

	queue = append(queue, start)
	for len(queue) != 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			current := queue[i]
			if string(current) == target {
				return step
			}
			// 找到'0'对应的index
			idx := 0
			for current[idx] != '0' {
				idx++
			}
			for _, x := range neighbor[idx] {
				tmp := make([]byte, 6)
				copy(tmp, current)
				tmp[idx], tmp[x] = tmp[x], tmp[idx]
				if _, ok := visited[string(tmp)]; !ok {
					queue = append(queue, tmp)
					visited[string(tmp)] = struct{}{}
				}
			}
		}
		queue = queue[sz:]
		step++
	}
	return -1
}
