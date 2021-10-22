// @file: max_heap.go
// @date: 2021/2/8

// Package heap
package heap

type maxHeap []int

func parent(x int) int {
	if x > 0 {
		return (x - 1) / 2
	}
	return 0
}

func right(x int) int {
	return 2*x + 2
}

func left(x int) int {
	return 2*x + 1
}

// 下沉i节点
func (h maxHeap) sink(i, n int) {
	// 保存要下沉的节点值
	tmp := h[i]
	for k := left(i); k < n; k = left(k) {
		// 如果有右孩子，则比较右孩子与左孩子的大小，取较大的那个
		if k+1 < n && h[k+1] > h[k] {
			k++
		}
		// 如果较大的孩子比要下沉的节点还要小，则不用继续下沉
		if h[k] <= tmp {
			break
		}
		// 将较大的孩子交换上来
		h[i] = h[k]
		i = k
	}
	// 将要下沉的节点值保存在交换后的空位
	h[i] = tmp
}

// 上浮i节点
func (h maxHeap) swim(i int) {
	for i > 1 && h[parent(i)] < h[i] {
		h[parent(i)], h[i] = h[i], h[parent(i)]
		i = parent(i)
	}
}

func (h *maxHeap) insert(x int) {
	*h = append(*h, x)
	h.swim(len(*h) - 1)
}

func (h *maxHeap) delMax() int {
	max := (*h)[0]
	n := len(*h)
	(*h)[0], (*h)[n-1] = (*h)[n-1], (*h)[0]
	*h = (*h)[:n-1]
	h.sink(0, n-1)
	return max
}
