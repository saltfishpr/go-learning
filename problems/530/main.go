package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	pre := 10000
	min := 10000

	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		now := abs(root.Val - pre)
		if now < min {
			min = now
		}
		pre = root.Val
		helper(root.Right)
	}

	helper(root)

	return min
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
