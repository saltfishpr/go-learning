package main

import "learning/data-structure/list"

type ListNode = list.ListNode[int]

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy ListNode
	tail := &dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}
		tail.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		tail = tail.Next
	}
	return dummy.Next
}
