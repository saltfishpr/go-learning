package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	cur := root
	for cur != nil {
		if cur.Child == nil {
			cur = cur.Next
		} else {
			childHead := flatten(cur.Child)
			childTail := tail(childHead)

			cur.Next, childTail.Next, childHead.Prev, cur, cur.Child = childHead, cur.Next, cur, cur.Next, nil
			if cur != nil {
				cur.Prev = childTail
			}
		}
	}
	return root
}

func tail(head *Node) *Node {
	for head.Next != nil {
		head = head.Next
	}
	return head
}
