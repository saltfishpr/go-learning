package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	idx := find(inorder, preorder[0])
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:idx+1], inorder[:idx]),
		Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
	}
}

func find(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}
