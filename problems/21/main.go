package main

import "learning/data-structure/list"

type ListNode = list.ListNode[int]

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	p := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1, list1.Next = list1.Next, nil
		} else {
			p.Next = list2
			list2, list2.Next = list2.Next, nil
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return dummy.Next
}
