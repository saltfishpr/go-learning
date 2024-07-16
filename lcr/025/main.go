package main

import "learning/data-structure/list"

type ListNode = list.ListNode[int]

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := list2Stack(l1), list2Stack(l2)
	carry := 0
	var res *ListNode
	for len(s1) != 0 || len(s2) != 0 || carry != 0 {
		var a, b int
		a, s1 = pop(s1)
		b, s2 = pop(s2)
		cur := a + b + carry
		carry = cur / 10
		cur %= 10
		node := &ListNode{Val: cur, Next: res}
		res = node
	}
	return res
}

func list2Stack(l *ListNode) []int {
	var res []int
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}

func pop(s []int) (int, []int) {
	if len(s) == 0 {
		return 0, nil
	}
	return s[len(s)-1], s[:len(s)-1]
}
