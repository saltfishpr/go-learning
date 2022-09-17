// @file: 32-2从上到下打印二叉树 II.go
// @date: 2021/2/18

// Package offer
package offer

/*
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
*/

func levelOrder2X32(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	res := make([][]int, 0)
	if root == nil {
		return res
	}
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
		res = append(res, tmp)
	}
	return res
}
