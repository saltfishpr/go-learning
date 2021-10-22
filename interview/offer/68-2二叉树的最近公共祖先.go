// @file: 68-2二叉树的最近公共祖先.go
// @date: 2021/2/26

// Package offer
package offer

func lowestCommonAncestorX68N2(root *TreeNode, p, q *TreeNode) *TreeNode {
	var helper func(root *TreeNode, p, q *TreeNode) *TreeNode
	helper = func(root *TreeNode, p, q *TreeNode) *TreeNode {
		if root == nil || root == p || root == q {
			return root
		}
		left := helper(root.Left, p, q)
		right := helper(root.Right, p, q)
		if left == nil {
			return right
		}
		return left
	}

	return helper(root, p, q)
}
