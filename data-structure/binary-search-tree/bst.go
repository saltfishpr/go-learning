// @file: bst.go
// @date: 2021/2/1

// Package bst
package bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 98. 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	res := make([]int, 0)
	var helper func(*TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		res = append(res, root.Val)
		helper(root.Right)
	}
	helper(root)
	if len(res) <= 1 {
		return true
	}
	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}
	return true
}

// 701. 二叉搜索树中的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	if root == nil {
		return node
	}
	p := root
	for {
		if p.Val >= val {
			if p.Left != nil {
				p = p.Left
			} else {
				p.Left = node
				break
			}
		} else {
			if p.Right != nil {
				p = p.Right
			} else {
				p.Right = node
				break
			}
		}
	}

	return root
}

func insertIntoBSTV2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	move, pre := root, root
	for move != nil {
		pre = move
		if val < move.Val {
			move = move.Left
		} else {
			move = move.Right
		}
	}
	if val > pre.Val {
		pre.Right = &TreeNode{Val: val}
	} else {
		pre.Left = &TreeNode{Val: val}
	}
	return root
}

// 230. 二叉搜索树中第K小的元素
func kthSmallest(root *TreeNode, k int) int {
	rank := 0
	res := 0

	var traverse func(*TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		rank++
		if rank == k {
			res = root.Val
			return
		} else {
			traverse(root.Right)
		}
	}

	traverse(root)
	return res
}

// 538. 把二叉搜索树转换为累加树
func convertBST(root *TreeNode) *TreeNode {
	sum := 0

	var traverse func(*TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Right)
		sum += root.Val
		root.Val = sum
		traverse(root.Left)
	}

	traverse(root)
	return root
}

// 1038. 把二叉搜索树转换为累加树
func bstToGst(root *TreeNode) *TreeNode {
	sum := 0

	var traverse func(*TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Right)
		sum += root.Val
		root.Val = sum
		traverse(root.Left)
	}

	traverse(root)
	return root
}

// 450. 删除二叉搜索树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key == root.Val {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		minNode := root.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		minNode.Left = root.Left
		return root.Right
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

// 700. 二叉搜索树中的搜索
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	for root != nil && root.Val != val {
		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}
