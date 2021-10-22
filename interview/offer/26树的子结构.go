// @file: 26树的子结构.go
// @date: 2021/2/15

// Package offer
package offer

/*
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
B是A的子结构， 即 A中有出现和B相同的结构和节点值。
*/

func isSubStructureX26(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	var helper func(A *TreeNode, B *TreeNode) bool
	helper = func(A *TreeNode, B *TreeNode) bool {
		if B == nil {
			return true
		}
		if A == nil {
			return false
		}
		return A.Val == B.Val && helper(A.Left, B.Left) && helper(A.Right, B.Right)
	}

	return helper(A, B) || isSubStructureX26(A.Left, B) || isSubStructureX26(A.Right, B)
}
