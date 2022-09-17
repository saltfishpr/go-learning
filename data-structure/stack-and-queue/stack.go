// @file: stack.go
// @date: 2021/1/23

// Package stackandqueue
package stackandqueue

import (
	"strconv"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type stackInt struct {
	length int
	stack  []int
}

func newStackInt() stackInt {
	return stackInt{
		length: 0,
		stack:  make([]int, 0),
	}
}

func (s *stackInt) push(x int) {
	s.stack = append(s.stack, x)
	s.length++
}

func (s *stackInt) pop() int {
	if s.length == 0 {
		return 0
	}
	res := s.top()
	s.stack = s.stack[:s.length-1]
	s.length--
	return res
}

func (s *stackInt) top() int {
	return s.stack[s.length-1]
}

// 150. 逆波兰表达式求值
func evalRPN(tokens []string) int {
	stack := newStackInt()
	for _, t := range tokens {
		switch t {
		case "+":
			b := stack.pop()
			a := stack.pop()
			stack.push(a + b)
		case "-":
			b := stack.pop()
			a := stack.pop()
			stack.push(a - b)
		case "*":
			b := stack.pop()
			a := stack.pop()
			stack.push(a * b)
		case "/":
			b := stack.pop()
			a := stack.pop()
			stack.push(a / b)
		default:
			x, _ := strconv.Atoi(t)
			stack.push(x)
		}
	}
	return stack.pop()
}

// 394. 字符串解码
func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}
	stack := make([]byte, 0)
	for i := range s {
		if s[i] == ']' {
			temp := make([]byte, 0)
			for len(stack) != 0 && stack[len(stack)-1] != '[' {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				temp = append(temp, v)
			}
			// pop '['
			stack = stack[:len(stack)-1]
			// pop num
			idx := 1
			for len(stack) >= idx && stack[len(stack)-idx] >= '0' && stack[len(stack)-idx] <= '9' {
				idx++
			}
			num := stack[len(stack)-idx+1:]
			count, _ := strconv.Atoi(string(num))
			stack = stack[:len(stack)-idx+1]
			for j := 0; j < count; j++ {
				for k := len(temp) - 1; k >= 0; k-- {
					stack = append(stack, temp[k])
				}
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

// 94. 二叉树的中序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root) // 压入
			root = root.Left
		}
		// 弹出
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		root = node.Right
	}
	return res
}

// 133. 克隆图
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node, 0)
	var clone func(*Node) *Node
	clone = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if v, ok := visited[node]; ok {
			return v
		}
		newNode := &Node{
			Val:       node.Val,
			Neighbors: make([]*Node, len(node.Neighbors)),
		}
		visited[node] = newNode
		for i := 0; i < len(node.Neighbors); i++ {
			newNode.Neighbors[i] = clone(node.Neighbors[i])
		}
		return newNode
	}

	return clone(node)
}

// 200. 岛屿数量
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	var count int
	row, col := len(grid), len(grid[0])

	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= row || y < 0 || y >= col {
			return 0
		}
		if grid[x][y] == '1' {
			grid[x][y] = '0'
			return dfs(x-1, y) + dfs(x+1, y) + dfs(x, y-1) + dfs(x, y+1) + 1
		}
		return 0
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' && dfs(i, j) > 0 {
				count++
			}
		}
	}

	return count
}

// 84. 柱状图中最大的矩形  单调栈法
func largestRectangleArea(heights []int) int {
	// 这里在最后添加一个0，可以将栈内元素全部弹出，简化代码
	heights = append(heights, 0)
	res := 0
	stack := newStackInt()

	for i := range heights {
		for stack.length > 0 && heights[i] < heights[stack.top()] {
			curHeight := heights[stack.pop()]
			// 把高度相同的柱视为一个柱
			for stack.length > 0 && curHeight == heights[stack.top()] {
				stack.pop()
			}
			// 计算宽度
			width := i
			if stack.length > 0 {
				width = i - stack.top() - 1
			}

			res = max(res, curHeight*width)
		}
		stack.push(i)
	}

	return res
}
