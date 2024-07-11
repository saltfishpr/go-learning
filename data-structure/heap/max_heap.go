package heap

import (
	"learning/pkg/constraints"
)

type MaxHeap[T constraints.Ordered] struct {
	data []T
}

func NewMaxHeap[T constraints.Ordered]() MaxHeap[T] {
	return MaxHeap[T]{}
}

func InitMaxHeap[T constraints.Ordered](values ...T) MaxHeap[T] {
	h := NewMaxHeap[T]()
	for _, v := range values {
		h.Push(v)
	}
	return h
}

// Push 元素入堆。
func (h *MaxHeap[T]) Push(val T) {
	// 添加节点
	h.data = append(h.data, val)
	// 从底至顶堆化
	h.siftUp(len(h.data) - 1)
}

// Pop 堆顶元素出堆。
func (h *MaxHeap[T]) Pop() T {
	if h.Size() == 0 {
		panic("heap is empty")
	}
	// 交换首元素与尾元素
	h.swap(0, h.Size()-1)
	// 删除节点
	val := h.data[h.Size()-1]
	h.data = h.data[:h.Size()-1]
	// 从顶至底堆化
	h.siftDown(0)
	// 返回堆顶元素
	return val
}

// Peek 访问堆顶元素。
func (h *MaxHeap[T]) Peek() T {
	return h.data[0]
}

// Size 返回堆的大小。
func (h *MaxHeap[T]) Size() int {
	return len(h.data)
}

// siftUp 从节点 i 开始，从底至顶堆化。
func (h *MaxHeap[T]) siftUp(i int) {
	for {
		// 获取节点 i 的父节点
		p := parent(i)
		// 当“越过根节点”或“节点无须修复”时，结束堆化
		if p < 0 || h.data[i] <= h.data[p] {
			break
		}
		// 交换两节点
		h.swap(i, p)
		// 循环向上堆化
		i = p
	}
}

// siftDown 从节点 i 开始，从顶至底堆化。
func (h *MaxHeap[T]) siftDown(i int) {
	for {
		// 判断节点 i, l, r 中值最大的节点，记为 max
		l, r, max := left(i), right(i), i
		if l < h.Size() && h.data[l] > h.data[max] {
			max = l
		}
		if r < h.Size() && h.data[r] > h.data[max] {
			max = r
		}
		// 若节点 i 最大或索引 l, r 越界，则无须继续堆化，跳出
		if max == i {
			break
		}
		// 交换最大值与当前节点
		h.swap(i, max)
		// 循环向下堆化
		i = max
	}
}

// swap 交换两个元素。
func (h *MaxHeap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// parent 返回节点 i 的父节点索引。
func parent(idx int) int {
	return (idx - 1) / 2
}

// left 返回节点 i 的左子节点索引。
func left(idx int) int {
	return idx*2 + 1
}

// right 返回节点 i 的右子节点索引。
func right(idx int) int {
	return idx*2 + 2
}
