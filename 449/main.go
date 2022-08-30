package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (*Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var preOrder []string

	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		preOrder = append(preOrder, strconv.Itoa(root.Val))
		helper(root.Left)
		helper(root.Right)
	}
	helper(root)

	return strings.Join(preOrder, ",")
}

// Deserializes your encoded data to tree.
func (*Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	ss := strings.Split(data, ",")
	preOrder := make([]int, len(ss))
	for i, s := range ss {
		val, _ := strconv.Atoi(s)
		preOrder[i] = val
	}

	var helper func(left, right int) *TreeNode
	helper = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		val := preOrder[left]
		mid := right + 1
		root := &TreeNode{val, nil, nil}
		for i := left + 1; i <= right; i++ {
			if preOrder[i] > val {
				mid = i
				break
			}
		}
		root.Left = helper(left+1, mid-1)
		root.Right = helper(mid, right)
		return root
	}

	return helper(0, len(preOrder)-1)
}
