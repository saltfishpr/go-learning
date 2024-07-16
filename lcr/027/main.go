package main

import "learning/data-structure/list"

type ListNode = list.ListNode[int]

func isPalindrome(head *ListNode) bool {
	mid := findMid(head)
	l1 := head
	l2 := reverseList(mid.Next)
	for l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}

// findMid 当链表长度为奇数时，返回中间节点；当链表长度为偶数时，返回中间靠前的节点。
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
