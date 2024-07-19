package main

import "learning/data-structure/tree"

type TreeNode = tree.TreeNode[int]

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}
