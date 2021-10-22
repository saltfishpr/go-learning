// @file: 25合并两个排序的链表.go
// @date: 2021/2/14

// Package offer
package offer

/*
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
*/

func mergeTwoListsX25(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	move := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			move.Next = l1
			l1 = l1.Next
			move = move.Next
		} else {
			move.Next = l2
			l2 = l2.Next
			move = move.Next
		}
	}
	for l1 != nil {
		move.Next = l1
		l1 = l1.Next
		move = move.Next
	}
	for l2 != nil {
		move.Next = l2
		l2 = l2.Next
		move = move.Next
	}
	return dummy.Next
}
