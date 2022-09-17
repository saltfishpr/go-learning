// @file: binary_tree.go
// @date: 2021/1/19

// Package binarytree
package binarytree

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

// 226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

// 116. 填充每个节点的下一个右侧节点指针
func connect(root *Node) *Node {
	var connectTwoNodes func(*Node, *Node)
	connectTwoNodes = func(node1, node2 *Node) {
		if node1 == nil || node2 == nil {
			return
		}
		node1.Next = node2
		connectTwoNodes(node1.Left, node1.Right)
		connectTwoNodes(node1.Right, node2.Left)
		connectTwoNodes(node2.Left, node2.Right)
	}
	if root == nil {
		return nil
	}
	connectTwoNodes(root.Left, root.Right)
	return root
}

// 114. 二叉树展开为链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	// 左右子树已经被拉成链表
	left, right := root.Left, root.Right
	// 将左子树接到右侧
	root.Left, root.Right = nil, left
	p := root
	for p.Right != nil {
		p = p.Right
	}
	// 右子树接到左子树后
	p.Right = right
}

// 124. 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	type resultType struct {
		singlePath int // 保存单边最大值
		maxPath    int // 保存最大值（单边或者两个单边+根的值）
	}

	var helper func(*TreeNode) resultType
	helper = func(root *TreeNode) resultType {
		if root == nil {
			return resultType{singlePath: 0, maxPath: math.MinInt64}
		}
		left := helper(root.Left)
		right := helper(root.Right)

		res := resultType{}
		if left.singlePath > right.singlePath {
			res.singlePath = max(left.singlePath+root.Val, 0)
		} else {
			res.singlePath = max(right.singlePath+root.Val, 0)
		}
		maxPath := max(left.maxPath, right.maxPath)
		res.maxPath = max(left.singlePath+right.singlePath+root.Val, maxPath)
		return res
	}

	return helper(root).maxPath
}

// 236. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

// 102. 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		tmp := make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tmp)
	}
	reverse(res)
	return res
}

func reverse(list [][]int) {
	length := len(list)
	for i := 0; i < length/2; i++ {
		list[i], list[length-1-i] = list[length-1-i], list[i]
	}
}

// 104. 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

// 110. 平衡二叉树
func isBalanced(root *TreeNode) bool {
	var helper func(*TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := helper(root.Left)
		right := helper(root.Right)
		if left == -1 || right == -1 || abs(left-right) > 1 {
			return -1
		}
		return max(left, right) + 1
	}
	if helper(root) == -1 {
		return false
	}
	return true
}

// 654. 最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}

	var build func([]int, int, int) *TreeNode
	build = func(nums []int, lo, hi int) *TreeNode {
		if lo > hi {
			return nil
		}
		idx, maxVal := -1, math.MinInt64
		for i := lo; i <= hi; i++ {
			if nums[i] > maxVal {
				idx, maxVal = i, nums[i]
			}
		}
		root := &TreeNode{Val: maxVal}
		root.Left = build(nums, lo, idx-1)
		root.Right = build(nums, idx+1, hi)
		return root
	}

	root := build(nums, 0, len(nums)-1)
	return root
}

// 105. 从前序与中序遍历序列构造二叉树
func buildTreePreAndIn(preorder []int, inorder []int) *TreeNode {
	preIdx := 0

	var build func([]int, int, int) *TreeNode
	build = func(inorder []int, lo int, hi int) *TreeNode {
		if lo > hi {
			return nil
		}
		rootIdx, rootVal := 0, preorder[preIdx]
		preIdx++
		for i := lo; i <= hi; i++ {
			if inorder[i] == rootVal {
				rootIdx = i
				break
			}
		}

		root := &TreeNode{Val: rootVal}
		root.Left = build(inorder, lo, rootIdx-1)
		root.Right = build(inorder, rootIdx+1, hi)
		return root
	}

	return build(inorder, 0, len(inorder)-1)
}

// 106. 从中序与后序遍历序列构造二叉树
func buildTreeInAndPost(inorder []int, postorder []int) *TreeNode {
	postIdx := len(postorder) - 1

	var build func([]int, int, int) *TreeNode
	build = func(inorder []int, lo int, hi int) *TreeNode {
		if lo > hi {
			return nil
		}
		rootIdx, rootVal := 0, postorder[postIdx]
		postIdx--
		for i := lo; i <= hi; i++ {
			if inorder[i] == rootVal {
				rootIdx = i
				break
			}
		}

		root := &TreeNode{Val: rootVal}
		root.Right = build(inorder, rootIdx+1, hi)
		root.Left = build(inorder, lo, rootIdx-1)
		return root
	}

	return build(inorder, 0, len(inorder)-1)
}

// 652. 寻找重复的子树
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m := make(map[string]int, 0)
	res := make([]*TreeNode, 0)

	var traverse func(*TreeNode) string
	traverse = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}
		left, right := traverse(root.Left), traverse(root.Right)
		total := fmt.Sprintf("%s,%s,%d", left, right, root.Val)
		if m[total] == 1 {
			res = append(res, root)
		}
		m[total]++
		return total
	}

	traverse(root)
	return res
}

// 297. 二叉树的序列化与反序列化
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	left := c.serialize(root.Left)
	right := c.serialize(root.Right)
	return fmt.Sprintf("%s,%s,%s", strconv.Itoa(root.Val), left, right)
}

func (c *Codec) deserialize(data string) *TreeNode {
	idx := 0
	nums := strings.Split(data, ",")

	var helper func([]string) *TreeNode
	helper = func(ss []string) *TreeNode {
		if idx >= len(ss) {
			return nil
		}
		s := ss[idx]
		idx++
		if s == "#" {
			return nil
		}
		val, _ := strconv.Atoi(s)
		node := &TreeNode{Val: val}
		node.Left = helper(ss)
		node.Right = helper(ss)
		return node
	}

	return helper(nums)
}
