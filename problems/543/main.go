package main

import "learning/data-structure/tree"

type TreeNode = tree.TreeNode[int]

func diameterOfBinaryTree(root *TreeNode) int {
	var res int

	var maxDepth func(*TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := maxDepth(root.Left)
		r := maxDepth(root.Right)

		res = max(res, l+r)

		return max(l, r) + 1
	}

	maxDepth(root)
	return res
}
