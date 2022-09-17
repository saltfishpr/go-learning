// @file: 55-2平衡二叉树.go
// @date: 2021/2/23

// Package offer
package offer

/*
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
*/

func isBalancedX52(root *TreeNode) bool {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	var helper func(*TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := helper(root.Left)
		right := helper(root.Right)
		if left == -1 || right == -1 || abs(right-left) > 1 {
			return -1
		}
		return max(left, right) + 1
	}

	return helper(root) != -1
}
