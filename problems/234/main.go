package main

import "learning/data-structure/list"

type ListNode = list.ListNode[int]

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	l1, l2 := head, reverse(slow)
	for l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1, l2 = l1.Next, l2.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev, current *ListNode = nil, head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
