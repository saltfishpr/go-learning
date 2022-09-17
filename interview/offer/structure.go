// @file: structure.go
// @date: 2021/2/12

// Package offer
package offer

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
