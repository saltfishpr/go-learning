package main

import "learning/data-structure/list"

type ListNode = list.ListNode[int]

func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	for head != nil {
		dummy.Next, head.Next, head = head, dummy.Next, head.Next
	}
	return dummy.Next
}
