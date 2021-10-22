// @file: queue.go
// @date: 2021/2/7

// Package monotonicqueue
package monotonicqueue

type node struct {
	val  int
	next *node
	prev *node
}

type queue struct {
	size int
	head *node
	tail *node
}

func newQueue() *queue {
	head := &node{val: 0}
	tail := &node{val: 0}
	head.next = tail
	tail.prev = head
	return &queue{size: 0, head: head, tail: tail}
}

func (q *queue) popLast() int {
	if q.size == 0 {
		return 0
	}
	x := q.tail.prev
	q.tail.prev = x.prev
	x.prev.next = q.tail
	q.size--
	return x.val
}

func (q *queue) popFirst() int {
	if q.size == 0 {
		return 0
	}
	x := q.head.next
	q.head.next = x.next
	x.next.prev = q.head
	q.size--
	return x.val
}

func (q *queue) append(val int) {
	x := &node{val: val}
	x.prev = q.tail.prev
	x.next = q.tail
	x.prev.next = x
	q.tail.prev = x
	q.size++
}

func (q *queue) getFirst() int {
	if q.size == 0 {
		return 0
	}
	return q.head.next.val
}

func (q *queue) isEmpty() bool {
	return q.size == 0
}

func (q *queue) push(n int) {
	for !q.isEmpty() && q.tail.prev.val < n {
		q.popLast()
	}
	q.append(n)
}

func (q *queue) max() int {
	return q.getFirst()
}

func (q *queue) pop(n int) {
	if n == q.getFirst() {
		q.popFirst()
	}
}

// 239. 滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	window := newQueue()
	res := make([]int, 0)
	for i := 0; i < k-1; i++ {
		window.push(nums[i])
	}
	for i := k - 1; i < len(nums); i++ {
		window.push(nums[i])
		res = append(res, window.max())
		window.pop(nums[i-k+1])
	}
	return res
}
