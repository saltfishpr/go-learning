// @file: 51数组中的逆序对.go
// @date: 2021/2/21

// Package offer
package offer

/*
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
*/

func reversePairsX51(nums []int) int {
	var mergeSort func([]int, int, int) int
	mergeSort = func(nums []int, left int, right int) int {
		if left >= right {
			return 0
		}

		mid := left + (right-left)/2
		cnt := mergeSort(nums, left, mid) + mergeSort(nums, mid+1, right)

		tmp := make([]int, 0)
		i, j := left, mid+1
		for i <= mid && j <= right {
			if nums[i] <= nums[j] {
				tmp = append(tmp, nums[i])
				cnt += j - (mid + 1) // 有 j-(mid+1) 个数比 num[i] 小
				i++
			} else {
				tmp = append(tmp, nums[j])
				j++
			}
		}
		for i <= mid {
			tmp = append(tmp, nums[i])
			cnt += right - (mid + 1) + 1 // 如果左边的数组有剩余，则每个数字的贡献度都为 right+1-(mid+1)
			i++
		}
		for j <= right {
			tmp = append(tmp, nums[j])
			j++
		}
		copy(nums[left:right+1], tmp)
		return cnt
	}

	return mergeSort(nums, 0, len(nums)-1)
}
