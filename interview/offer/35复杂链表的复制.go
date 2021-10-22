// @file: 35复杂链表的复制.go
// @date: 2021/2/18

// Package offer
package offer

/*
请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。
*/

func copyRandomListX35(head *Node) *Node {
	oldToNew := make(map[*Node]*Node, 0)
	var cur *Node

	cur = head
	for cur != nil {
		oldToNew[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		oldToNew[cur].Next = oldToNew[cur.Next]
		oldToNew[cur].Random = oldToNew[cur.Random]
		cur = cur.Next
	}

	return oldToNew[head]
}

func copyRandomListX35V2(head *Node) *Node {
	if head == nil {
		return nil
	}
	var cur *Node

	// 将复制的节点接在旧节点后
	cur = head
	for cur != nil {
		clone := &Node{Val: cur.Val, Next: cur.Next, Random: nil}
		cur.Next, cur = clone, clone.Next
	}

	// 处理Random
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// 分离新旧链表
	cur = head
	curHead := head.Next
	for cur != nil && cur.Next != nil {
		cur.Next, cur = cur.Next.Next, cur.Next
	}

	return curHead
}
