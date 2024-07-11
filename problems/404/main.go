package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res int

	if isLeaf(root.Left) {
		res += root.Left.Val
	} else {
		res += sumOfLeftLeaves(root.Left)
	}

	res += sumOfLeftLeaves(root.Right)

	return res
}

func isLeaf(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return root.Left == nil && root.Right == nil
}
