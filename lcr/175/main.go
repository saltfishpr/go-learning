package main

import "learning/data-structure/tree"

type TreeNode = tree.TreeNode[int]

func calculateDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := calculateDepth(root.Left)
	r := calculateDepth(root.Right)
	return max(l, r) + 1
}
