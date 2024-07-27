package main

import "learning/data-structure/list"

type ListNode = list.ListNode[int]

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverse(head *ListNode, tail *ListNode) *ListNode {
	var prev, current, next *ListNode = nil, head, head
	for current != tail {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
