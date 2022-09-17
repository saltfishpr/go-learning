// @file: 33二叉搜索树的后序遍历序列.go
// @date: 2021/2/18

// Package offer
package offer

/*
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同
*/

func verifyPostorderX33(postorder []int) bool {
	var helper func(int, int) bool
	helper = func(i, j int) bool {
		if i >= j {
			return true
		}
		root := postorder[j]
		p := i
		for postorder[p] < root {
			p++
		}
		m := p
		for postorder[p] > root {
			p++
		}
		return p == j && helper(i, m-1) && helper(m, j-1)
	}

	return helper(0, len(postorder)-1)
}
