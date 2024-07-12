package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	for head != nil {
		dummy.Next, head.Next, head = head, dummy.Next, head.Next
	}
	return dummy.Next
}
