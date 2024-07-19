package main

import "learning/data-structure/tree"

type TreeNode = tree.TreeNode[int]

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	return max(l, r) + 1
}
