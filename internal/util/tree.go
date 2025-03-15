package util

type TreeNode[T any] struct {
	Val      T
	Children []*TreeNode[T]
}

// AddChild 添加子节点
func (n *TreeNode[T]) AddChild(child *TreeNode[T]) {
	n.Children = append(n.Children, child)
}

// DFS 深度优先搜索
func (n *TreeNode[T]) DFS(target T, equals func(a, b T) bool) *TreeNode[T] {
	if equals(n.Val, target) {
		return n
	}
	for _, child := range n.Children {
		if result := child.DFS(target, equals); result != nil {
			return result
		}
	}
	return nil
}
