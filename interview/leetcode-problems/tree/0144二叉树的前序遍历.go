// @file: 0144二叉树的前序遍历.go
// @description: 给定一个二叉树，返回它的前序遍历。
// @author: SaltFish
// @date: 2020/09/09

package tree

/*
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

func preorderTraversal(root *TreeNode) []int {
	var helper func(*TreeNode)
	var res []int
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		helper(root.Left)
		helper(root.Right)
	}
	helper(root)
	return res
}

func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var stack []*TreeNode
	var output []int
	stack = append(stack, root)
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root != nil {
			output = append(output, root.Val)
			if root.Right != nil { // 先入后出
				stack = append(stack, root.Right)
			}
			if root.Left != nil {
				stack = append(stack, root.Left)
			}
		}
	}
	return output
}

func preorderTraversal3(root *TreeNode) []int {
	node := root
	var output []int
	for node != nil {
		if node.Left == nil {
			output = append(output, node.Val)
			node = node.Right
		} else {
			predecessor := node.Left
			for predecessor.Right != nil && predecessor.Right != node {
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				output = append(output, node.Val)
				predecessor.Right = node
				node = node.Left
			} else {
				predecessor.Right = nil
				node = node.Right
			}
		}
	}
	return output
}
