// @file: 28对称的二叉树.go
// @date: 2021/2/15

// Package offer
package offer

/*
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
*/

func isSymmetricX28(root *TreeNode) bool {
	var helper func(*TreeNode, *TreeNode) bool
	helper = func(A *TreeNode, B *TreeNode) bool {
		if A != nil && B != nil {
			return A.Val == B.Val && helper(A.Left, B.Right) && helper(A.Right, B.Left)
		}
		return A == nil && B == nil
	}

	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}
