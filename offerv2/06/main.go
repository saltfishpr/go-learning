package main

import "learning/data-structure/list"

type ListNode = list.ListNode[int]

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var res []int
	var helper func(*ListNode)
	helper = func(node *ListNode) {
		if node.Next != nil {
			helper(node.Next)
		}
		res = append(res, node.Val)
	}

	helper(head)
	return res
}
