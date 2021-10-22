// @file: offer41.go
// @date: 2021/2/18

// Package offer41
package offer41

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	large IntHeap // 小顶堆，储存较大的值
	small IntHeap // 大顶堆，储存较小的值
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		large: make([]int, 0),
		small: make([]int, 0),
	}
}

func (f *MedianFinder) AddNum(num int) {
	if len(f.large) >= len(f.small) {
		heap.Push(&f.large, num)
		heap.Push(&f.small, -heap.Pop(&f.large).(int))
	} else {
		heap.Push(&f.small, -num)
		heap.Push(&f.large, -heap.Pop(&f.small).(int))
	}
}

func (f *MedianFinder) FindMedian() float64 {
	if len(f.small) == len(f.large) {
		return float64(-f.small[0]+f.large[0]) / 2
	}
	return float64(-f.small[0])
}
