package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	var res int

	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := helper(root.Left)
		right := helper(root.Right)
		var left_res, right_res int
		if root.Left != nil && root.Left.Val == root.Val {
			left_res = left + 1
		}
		if root.Right != nil && root.Right.Val == root.Val {
			right_res = right + 1
		}
		res = max(res, left_res+right_res)
		return max(left_res, right_res)
	}

	helper(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
