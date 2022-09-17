// @file: 68-1二叉搜索树的最近公共祖先.go
// @date: 2021/2/26

// Package offer
package offer

func lowestCommonAncestorX68N1(root *TreeNode, p, q *TreeNode) *TreeNode {
	// 保证 p 的值较小
	if p.Val > q.Val {
		p, q = q, p
	}
	for root != nil {
		if root.Val < p.Val {
			root = root.Right
		} else if root.Val > q.Val {
			root = root.Left
		} else {
			break
		}
	}
	return root
}
