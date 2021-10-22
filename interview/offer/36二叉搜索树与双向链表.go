// @file: 36二叉搜索树与双向链表.go
// @date: 2021/2/18

// Package offer
package offer

/*
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。
*/

func treeToDoublyListX36(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var pre, head *TreeNode

	var helper func(*TreeNode)
	helper = func(cur *TreeNode) {
		helper(cur.Left)
		if pre != nil {
			pre.Right, cur.Left = cur, pre
		} else {
			head = cur
		}
		pre = cur
		helper(cur.Right)
	}

	helper(root)
	head.Left, pre.Right = pre, head
	return head
}
