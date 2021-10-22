// @file: 34二叉树中和为某一值的路径.go
// @date: 2021/2/18

// Package offer
package offer

/*
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
*/

func pathSumX34(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var helper func(*TreeNode, int)
	helper = func(node *TreeNode, tar int) {
		if node == nil {
			return
		}
		tar -= node.Val
		path = append(path, node.Val)

		if tar == 0 && node.Left == nil && node.Right == nil {
			// 复制path
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}

		helper(node.Left, tar)
		helper(node.Right, tar)

		path = path[:len(path)-1]
	}

	helper(root, sum)
	return res
}
