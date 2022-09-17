// @file: min_stack.go
// @date: 2021/1/23

// Package minstack
package minstack

import "math"

// 155. 最小栈
type MinStack struct {
	min   []int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{min: make([]int, 0), stack: make([]int, 0)}
}

func (s *MinStack) Push(x int) {
	m := s.GetMin()
	if x < m {
		s.min = append(s.min, x)
	} else {
		s.min = append(s.min, m)
	}
	s.stack = append(s.stack, x)
}

func (s *MinStack) Pop() {
	if len(s.stack) == 0 {
		return
	}
	s.stack = s.stack[:len(s.stack)-1]
	s.min = s.min[:len(s.min)-1]
}

func (s *MinStack) Top() int {
	if len(s.stack) == 0 {
		return 0
	}
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	if len(s.min) == 0 {
		return math.MaxInt64
	}
	return s.min[len(s.min)-1]
}
