package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		max := queue[0].Val
		var nextQueue []*TreeNode
		for _, node := range queue {
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		res = append(res, max)
		queue = nextQueue
	}

	return res
}
