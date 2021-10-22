// @file: 32-1从上到下打印二叉树.go
// @date: 2021/2/18

// Package offer
package offer

/*
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
*/

func levelOrderX32(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}
