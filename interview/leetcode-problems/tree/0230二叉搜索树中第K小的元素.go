// @file: 0230二叉搜索树中第K小的元素.go
// @description: 给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。
// @author: SaltFish
// @date: 2020/09/10

package tree

/*
说明：
你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。

示例 1:
输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 1

示例 2:
输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 3
进阶：
如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？
*/

func kthSmallest(root *TreeNode, k int) int {
	var helper func(*TreeNode)
	var res []int
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		res = append(res, root.Val)
		helper(root.Right)
	}
	helper(root)
	return res[k-1]
}

func kthSmallest2(root *TreeNode, k int) int {
	var stack []*TreeNode
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if k--; k == 0 {
			return root.Val
		}
		root = root.Right
	}
}
