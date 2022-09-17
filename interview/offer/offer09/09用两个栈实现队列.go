// @file: offer09.go
// @date: 2021/2/12

// Package offer09
package offer09

/*
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

示例 1：
输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]

示例 2：
输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
*/

func push(nums *[]int, x int) {
	*nums = append(*nums, x)
}

func pop(nums *[]int) {
	n := len(*nums)
	if n > 0 {
		*nums = (*nums)[:n-1]
	}
}

func peek(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	return nums[n-1]
}

type CQueue struct {
	head, tail []int
}

func Constructor() CQueue {
	return CQueue{
		head: make([]int, 0),
		tail: make([]int, 0),
	}
}

func (q *CQueue) AppendTail(value int) {
	push(&q.tail, value)
}

func (q *CQueue) DeleteHead() int {
	var res int
	if len(q.head) == 0 {
		if len(q.tail) == 0 {
			return -1
		}
		for len(q.tail) != 0 {
			push(&q.head, peek(q.tail))
			pop(&q.tail)
		}
	}
	res = peek(q.head)
	pop(&q.head)
	return res
}
