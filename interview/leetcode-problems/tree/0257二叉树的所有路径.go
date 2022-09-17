// @file: 0257二叉树的所有路径.go
// @description: 给定一个二叉树，返回所有从根节点到叶子节点的路径。
// @author: SaltFish
// @date: 2020/09/10

package tree

/*
说明: 叶子节点是指没有子节点的节点。

示例:
输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
*/

import (
	"bytes"
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	var dfs func(*TreeNode, []int)
	var helper func([]int) string
	dfs = func(node *TreeNode, path []int) {
		if node == nil {
			return
		}
		pathB := append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			paths = append(paths, helper(pathB))
		}
		dfs(node.Left, pathB)
		dfs(node.Right, pathB)
	}
	helper = func(nums []int) string {
		var buffer bytes.Buffer
		for _, x := range nums {
			buffer.WriteString(strconv.Itoa(x) + "->")
		}
		res := buffer.String()
		return res[:len(res)-2]
	}
	dfs(root, []int{})
	return paths
}

func binaryTreePaths1(root *TreeNode) []string {
	var paths []string
	var constructPaths func(*TreeNode, string)
	constructPaths = func(root *TreeNode, path string) {
		if root != nil {
			pathSB := path
			pathSB += strconv.Itoa(root.Val)
			if root.Left == nil && root.Right == nil {
				paths = append(paths, pathSB)
			} else {
				pathSB += "->"
				constructPaths(root.Left, pathSB)
				constructPaths(root.Right, pathSB)
			}
		}
	}
	constructPaths(root, "")
	return paths
}
