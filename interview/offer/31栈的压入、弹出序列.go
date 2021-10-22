// @file: 31栈的压入、弹出序列.go
// @date: 2021/2/18

// Package offer
package offer

/*
输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。
假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。
*/

func validateStackSequencesX31(pushed []int, popped []int) bool {
	i, j := 0, 0
	n := len(pushed)
	stack := make([]int, 0)
	for i < n {
		stack = append(stack, pushed[i])
		i++
		for len(stack) != 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}

	return len(stack) == 0
}
