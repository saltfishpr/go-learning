// @file: median_finder.go
// @date: 2021/2/6

// Package heap
package heap

// 295. 数据流的中位数
type MedianFinder struct {
	large []int // 小顶堆，储存较大的值
	small []int // 大顶堆，储存较小的值
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		large: make([]int, 0),
		small: make([]int, 0),
	}
}

func (f *MedianFinder) AddNum(num int) {
	if len(f.small) >= len(f.large) {
		f.small = insert(f.small, num, true)
		f.large = insert(f.large, f.small[0], false)
		f.small = f.small[1:]
	} else {
		f.large = insert(f.large, num, false)
		f.small = insert(f.small, f.large[0], true)
		f.large = f.large[1:]
	}
}

func (f *MedianFinder) FindMedian() float64 {
	if len(f.small) == len(f.large) {
		return float64(f.small[0]+f.large[0]) / 2
	}
	if len(f.small) > len(f.large) {
		return float64(f.small[0])
	}
	return float64(f.large[0])
}

func insert(nums []int, value int, reverse bool) []int {
	n := len(nums)
	if n == 0 {
		nums = append(nums, value)
		return nums
	}
	res := make([]int, n+1)
	for i := 0; i < n; i++ {
		if !reverse {
			if nums[i] > value {
				copy(res[:i], nums[:i])
				res[i] = value
				copy(res[i+1:], nums[i:])
				return res
			}
		} else {
			if nums[i] < value {
				copy(res[:i], nums[:i])
				res[i] = value
				copy(res[i+1:], nums[i:])
				return res
			}
		}
	}
	res[n] = value
	copy(res[:n], nums)
	return res
}
