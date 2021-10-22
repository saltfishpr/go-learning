// @file: 0145二叉树的后序遍历.go
// @description: 给定一个二叉树，返回它的 后序 遍历。
// @author: SaltFish
// @date: 2020/09/10

package tree

/*
示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
*/

func postorderTraversal(root *TreeNode) []int {
	var res []int
	var helper func(*TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		helper(root.Right)
		res = append(res, root.Val)
	}
	helper(root)
	return res
}
