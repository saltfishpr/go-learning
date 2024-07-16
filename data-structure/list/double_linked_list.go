package list

type DoubleLinkedListNode[K comparable, V any] struct {
	Val  V
	Prev *DoubleLinkedListNode[K, V]
	Next *DoubleLinkedListNode[K, V]
}

type DoubleLinkedList[K comparable, V any] struct {
	dummyHead *DoubleLinkedListNode[K, V]
	dummyTail *DoubleLinkedListNode[K, V]
}

func NewDoubleLinkedList[K comparable, V any]() *DoubleLinkedList[K, V] {
	dummyHead := &DoubleLinkedListNode[K, V]{}
	dummyTail := &DoubleLinkedListNode[K, V]{}
	dummyHead.Next = dummyTail
	dummyTail.Prev = dummyHead
	return &DoubleLinkedList[K, V]{
		dummyHead: dummyHead,
		dummyTail: dummyTail,
	}
}

func (l *DoubleLinkedList[K, V]) Insert(idx int, node *DoubleLinkedListNode[K, V]) {
	if idx < 0 {
		return
	}

	current := l.dummyHead
	for i := 0; i < idx; i++ {
		if current.Next == l.dummyTail {
			break
		}
		current = current.Next
	}
	current.Next, current.Next.Prev, node.Next, node.Prev = node, node, current.Next, current
}

func (l *DoubleLinkedList[K, V]) Remove(node *DoubleLinkedListNode[K, V]) {
	if node == nil || node == l.dummyHead || node == l.dummyTail {
		return
	}
	node.Prev.Next, node.Next.Prev, node.Prev, node.Next = node.Next, node.Prev, nil, nil
}
