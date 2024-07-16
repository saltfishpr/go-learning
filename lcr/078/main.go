package main

import (
	"container/heap"

	"learning/data-structure/list"
)

type ListNode = list.ListNode[int]

func mergeKLists(lists []*ListNode) *ListNode {
	pq := &PriorityQueue{}
	for _, l := range lists {
		if l != nil {
			heap.Push(pq, l)
		}
	}
	dummy := &ListNode{}
	p := dummy
	for pq.Len() != 0 {
		node := heap.Pop(pq).(*ListNode)
		p.Next = node
		p = p.Next
		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
	}
	return dummy.Next
}

type PriorityQueue struct {
	lists []*ListNode
}

func (p *PriorityQueue) Len() int {
	return len(p.lists)
}

func (p *PriorityQueue) Less(i int, j int) bool {
	return p.lists[i].Val < p.lists[j].Val
}

func (p *PriorityQueue) Pop() any {
	n := p.Len()
	node := p.lists[n-1]
	p.lists = p.lists[:n-1]
	return node
}

func (p *PriorityQueue) Push(x any) {
	node := x.(*ListNode)
	p.lists = append(p.lists, node)
}

func (p *PriorityQueue) Swap(i int, j int) {
	p.lists[i], p.lists[j] = p.lists[j], p.lists[i]
}
