package main

import "learning/data-structure/list"

type ListNode = list.ListNode[int]

func reorderList(head *ListNode) {
	l1, l2 := head, findMid(head)
	// 分离两个链表
	l2.Next, l2 = nil, l2.Next
	// 反转第二个链表
	l2 = reverseList(l2)

	for l1 != nil && l2 != nil {
		l1.Next, l2.Next, l1, l2 = l2, l1.Next, l1.Next, l2.Next
	}
}

func findMid(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	for head != nil {
		dummy.Next, head.Next, head = head, dummy.Next, head.Next
	}
	return dummy.Next
}
