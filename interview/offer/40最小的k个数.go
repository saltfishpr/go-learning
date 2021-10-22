// @file: 40最小的k个数.go
// @date: 2021/2/18

// Package offer
package offer

/*
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
*/

func getLeastNumbersX40(arr []int, k int) []int {
	partition := func(nums []int, l int, r int) int {
		if l >= r {
			return l
		}
		pivot := nums[r]
		i := l - 1
		for j := l; j < r; j++ {
			if nums[j] < pivot {
				i++
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		nums[i+1], nums[r] = nums[r], nums[i+1]
		// 排序结果：第 i+1 位前的数都小于第 i+1 位
		return i + 1
	}
	left, right := 0, len(arr)-1
	pos := partition(arr, left, right)
	for pos != k {
		if k > pos {
			left = pos + 1
		} else if k < pos {
			right = pos - 1
		}
		pos = partition(arr, left, right)
	}

	return arr[:k]
}
