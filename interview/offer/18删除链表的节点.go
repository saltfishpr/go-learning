// @file: 18删除链表的节点.go
// @date: 2021/2/13

// Package offer
package offer

/*
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。

返回删除后的链表的头节点。

注意：此题对比原题有改动

示例 1:
输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
*/

func deleteNodeX18(head *ListNode, val int) *ListNode {
	newHead := &ListNode{Val: 0}
	newHead.Next = head
	pre := newHead
	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
			break
		}
		pre = head
		head = head.Next
	}
	return newHead.Next
}

func deleteNodeX18V2(head *ListNode, val int) *ListNode {
	move := head
	if head.Val == val {
		return head.Next
	}
	for move.Next != nil {
		if move.Next.Val == val {
			move.Next = move.Next.Next
			return head
		}
		move = move.Next
	}
	return head
}
