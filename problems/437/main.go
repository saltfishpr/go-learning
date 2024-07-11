package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) (ans int) {
	preSum := map[int]int{0: 1}

	var helper func(*TreeNode, int)
	// cur 当前的前缀和
	helper = func(node *TreeNode, cur int) {
		if node == nil {
			return
		}
		cur += node.Val
		// 是否存在前缀和刚好等于 cur−targetSum
		ans += preSum[cur-targetSum]
		preSum[cur]++
		helper(node.Left, cur)
		helper(node.Right, cur)
		// 回溯
		preSum[cur]--
	}
	helper(root, 0)

	return
}
