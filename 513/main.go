package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var res int

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		res = queue[0].Val
		var nextQueue []*TreeNode
		for _, node := range queue {
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		if len(nextQueue) == 0 {
			break
		}
		queue = nextQueue
	}

	return res
}
