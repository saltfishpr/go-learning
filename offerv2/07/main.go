package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	idx := find(inorder, preorder[0])
	root := &TreeNode{
		Val:   preorder[0],
		Left:  &TreeNode{},
		Right: &TreeNode{},
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
