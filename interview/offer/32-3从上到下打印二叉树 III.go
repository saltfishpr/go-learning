// @file: 32-3从上到下打印二叉树 III.go
// @date: 2021/2/18

// Package offer
package offer

/*
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
*/

func levelOrder3X32(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	reverse := func(nums []int, i, j int) {
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	queue := make([]*TreeNode, 0)
	re := false
	queue = append(queue, root)
	for len(queue) != 0 {
		n := len(queue)
		tmp := make([]int, 0)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if re {
			reverse(tmp, 0, len(tmp)-1)
		}
		re = !re
		res = append(res, tmp)
	}
	return res
}
