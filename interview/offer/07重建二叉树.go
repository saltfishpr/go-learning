// @file: 07重建二叉树.go
// @date: 2021/2/12

// Package offer
package offer

/*
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

例如，给出
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]

返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
*/

func buildTreeX07(preorder []int, inorder []int) *TreeNode {
	pos := 0

	var helper func([]int, []int) *TreeNode
	helper = func(preorder []int, inorder []int) *TreeNode {
		if len(inorder) == 0 {
			return nil
		}
		val := preorder[pos]
		pos++

		idx := 0
		for inorder[idx] != val {
			idx++
		}

		root := &TreeNode{Val: val}
		root.Left = helper(preorder, inorder[:idx])
		root.Right = helper(preorder, inorder[idx+1:])
		return root
	}

	return helper(preorder, inorder)
}
