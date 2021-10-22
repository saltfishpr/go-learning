// @file: 24反转链表.go
// @date: 2021/2/14

// Package offer
package offer

/*
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
*/

func reverseListX24(head *ListNode) *ListNode {
	var pre, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
