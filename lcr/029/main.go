package main

type Node struct {
	Val  int
	Next *Node
}

func insert(head *Node, x int) *Node {
	node := &Node{Val: x}
	if head == nil {
		node.Next = node
		return node
	}
	if head.Next == head {
		head.Next = node
		node.Next = head
		return head
	}
	curr, next := head, head.Next
	for next != head {
		if next.Val >= x && x >= curr.Val {
			break
		}
		if curr.Val > next.Val {
			if x > curr.Val || x < next.Val {
				break
			}
		}
		curr = curr.Next
		next = next.Next
	}
	curr.Next = node
	node.Next = next
	return head
}
