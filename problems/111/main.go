package main

import "learning/data-structure/tree"

type TreeNode = tree.TreeNode[int]

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		depth++
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			node := queue[i]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[queueSize:]
	}
	return -1
}
