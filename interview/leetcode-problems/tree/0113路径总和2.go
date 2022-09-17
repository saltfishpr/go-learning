// @file: 0113路径总和2.go
// @description: 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
// @author: SaltFish
// @date: 2020/09/02

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
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]

*/

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	path := []int{}
	if root == nil {
		return res
	}
	dfs(&res, root, path, sum)
	return res
}

func dfs(ret *[][]int, root *TreeNode, path []int, target int) {
	switch {
	case root == nil:
		return
	case root.Left == nil && root.Right == nil && target == root.Val:
		dst := []int{}
		copy(dst, append(path, root.Val))
		*ret = append(*ret, dst)
		return
	}
	path = append(path, root.Val)
	dfs(ret, root.Left, path, target-root.Val)
	dfs(ret, root.Right, path, target-root.Val)
}
