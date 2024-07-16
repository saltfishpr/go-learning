package list

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

func NewLinkedList[T any](values ...T) *ListNode[T] {
	dummy := &ListNode[T]{}
	current := dummy

	for _, val := range values {
		current.Next = &ListNode[T]{Val: val}
		current = current.Next
	}

	return dummy.Next
}
