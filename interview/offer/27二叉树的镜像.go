// @file: 27二叉树的镜像.go
// @date: 2021/2/15

// Package offer
package offer

/*
请完成一个函数，输入一个二叉树，该函数输出它的镜像。
*/

func mirrorTreeX27(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := mirrorTreeX27(root.Left)
	right := mirrorTreeX27(root.Right)
	root.Left = right
	root.Right = left
	return root
}
