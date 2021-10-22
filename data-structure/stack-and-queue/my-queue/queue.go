// @file: queue.go
// @date: 2021/1/23

// Package myqueue
package myqueue

// 232. 用栈实现队列
type MyQueue struct {
	head []int
	tail []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		head: make([]int, 0),
		tail: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	q.tail = append(q.tail, x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	if q.Empty() {
		return 0
	}
	var res int
	if len(q.head) > 0 {
		res = q.head[len(q.head)-1]
		q.head = q.head[:len(q.head)-1]
		return res
	}
	for len(q.tail) > 0 {
		q.head = append(q.head, q.tail[len(q.tail)-1])
		q.tail = q.tail[:len(q.tail)-1]
	}
	res = q.head[len(q.head)-1]
	q.head = q.head[:len(q.head)-1]
	return res
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	if len(q.head) > 0 {
		return q.head[len(q.head)-1]
	}
	for len(q.tail) > 0 {
		q.head = append(q.head, q.tail[len(q.tail)-1])
		q.tail = q.tail[:len(q.tail)-1]
	}
	return q.head[len(q.head)-1]
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	if len(q.head) == 0 && len(q.tail) == 0 {
		return true
	}
	return false
}
