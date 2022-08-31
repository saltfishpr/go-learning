package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	m := make(map[int]int)

	var helper func(*TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		val := helper(root.Left) + helper(root.Right) + root.Val
		m[val]++
		return val
	}

	helper(root)

	var maxCount int
	for _, v := range m {
		if v > maxCount {
			maxCount = v
		}
	}

	var res []int
	for k, v := range m {
		if v == maxCount {
			res = append(res, k)
		}
	}

	return res
}
