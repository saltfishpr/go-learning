// @file: offer37.go
// @date: 2021/2/18

// Package offer37
package offer37

import (
	"fmt"
	"strconv"
	"strings"
)

/*
请实现两个函数，分别用来序列化和反序列化二叉树。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	left := c.serialize(root.Left)
	right := c.serialize(root.Right)
	return fmt.Sprintf("%d,%s,%s", root.Val, left, right)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	nums := strings.Split(data, ",")
	n := len(nums)
	idx := 0

	var helper func() *TreeNode
	helper = func() *TreeNode {
		if idx >= n {
			return nil
		}
		s := nums[idx]
		idx++
		if s == "#" {
			return nil
		}
		val, _ := strconv.Atoi(s)
		node := &TreeNode{Val: val}
		node.Left = helper()
		node.Right = helper()
		return node
	}

	return helper()
}
