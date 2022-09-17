// @file: 54二叉搜索树的第k大节点.go
// @date: 2021/2/23

// Package offer
package offer

/*
给定一棵二叉搜索树，请找出其中第k大的节点。
*/

func kthLargestX54(root *TreeNode, k int) int {
	res := 0

	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Right)
		if k == 0 {
			return
		}
		k--
		if k == 0 {
			res = node.Val
		}
		helper(node.Left)
	}

	helper(root)
	return res
}
