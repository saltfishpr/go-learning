package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var queue []*Node
	queue = append(queue, root)
	var res [][]int
	for len(queue) != 0 {
		length := len(queue)
		var temp []int
		for i := 0; i < length; i++ {
			node := queue[i]
			temp = append(temp, node.Val)
			queue = append(queue, node.Children...)
		}
		res = append(res, temp)
		queue = queue[length:]
	}
	return res
}
