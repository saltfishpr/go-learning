// @file: 0112路径总和1.go
// @description: 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
// @author: SaltFish
// @date: 2020/08/26

package tree

/*
说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。

DFS
*/

// 递归
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

// 广度优先搜索
func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	queNode := []*TreeNode{}
	queVal := []int{}
	queNode = append(queNode, root)
	queVal = append(queVal, root.Val)
	for len(queNode) != 0 {
		node := queNode[0]
		queNode = queNode[1:]
		temp := queVal[0]
		queVal = queVal[1:]
		if node.Left == nil && node.Right == nil {
			if temp == sum {
				return true
			}
			continue
		}
		if node.Left != nil {
			queNode = append(queNode, node.Left)
			queVal = append(queVal, node.Left.Val+temp)
		}
		if node.Right != nil {
			queNode = append(queNode, node.Right)
			queVal = append(queVal, node.Right.Val+temp)
		}
	}
	return false
}
