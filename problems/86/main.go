package main

func main() {
	head := sliceToList(1, 4, 3, 2, 5, 2)
	partition(head, 3)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{}
	dummy2 := &ListNode{}
	p1 := dummy1
	p2 := dummy2

	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}
		head, head.Next = head.Next, nil
	}

	p1.Next = dummy2.Next
	return dummy1.Next
}

func sliceToList(slice ...int) *ListNode {
	if len(slice) == 0 {
		return nil
	}

	// 初始化链表头节点
	head := &ListNode{Val: slice[0]}
	current := head

	// 迭代 slice，创建链表
	for _, val := range slice[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head
}
