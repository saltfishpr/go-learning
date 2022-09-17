// @file: 0297二叉树的序列化与反序列化.go
// @description:
// @author: SaltFish
// @date: 2020/09/10

package tree

/*
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

示例:
你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"

提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

说明: 不要使用类的成员/全局/静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。
*/

import (
	"strconv"
	"strings"
)

type Codec struct {
	l []string
}

// func Constructor() Codec {
// 	return Codec{}
// }

func rserialize(root *TreeNode, str string) string {
	if root == nil {
		str += "null,"
	} else {
		str += strconv.Itoa(root.Val) + ","
		str = rserialize(root.Left, str)
		str = rserialize(root.Right, str)
	}
	return str
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	return rserialize(root, "")
}

func (c *Codec) rdeserialize() *TreeNode {
	if c.l[0] == "null" {
		c.l = c.l[1:]
		return nil
	}

	val, _ := strconv.Atoi(c.l[0])
	root := &TreeNode{Val: val}
	c.l = c.l[1:]
	root.Left = c.rdeserialize()
	root.Right = c.rdeserialize()
	return root
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	l := strings.Split(data, ",")
	for i := 0; i < len(l); i++ {
		if l[i] != "" {
			c.l = append(c.l, l[i])
		}
	}
	return c.rdeserialize()
}
