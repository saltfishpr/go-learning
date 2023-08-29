package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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
