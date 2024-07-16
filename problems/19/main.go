package main

import "learning/data-structure/list"

type ListNode = list.ListNode[int]

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p := getNthFromEnd(dummy, n+1) // 找到 n 的前置节点
	p.Next = p.Next.Next
	return dummy.Next
}

func getNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
