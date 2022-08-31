package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	m := make(map[int]int)
	var mode int

	var helper func(*TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		m[root.Val]++
		if m[root.Val] > mode {
			mode = m[root.Val]
		}
		helper(root.Right)
	}

	helper(root)

	var res []int
	for k, v := range m {
		if v == mode {
			res = append(res, k)
		}
	}

	return res
}
