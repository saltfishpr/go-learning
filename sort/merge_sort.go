// @file: merge_sort.go
// @description: 归并排序
// @author: SaltFish
// @date: 2020/9/12

package sort

func merge(left, right []int) (res []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
		}
	}
	res = append(res, left[l:]...) // ...打散数组
	res = append(res, right[r:]...)
	return
}

func mergeSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	m := length / 2
	left := mergeSort(nums[:m])
	right := mergeSort(nums[m:])
	return merge(left, right)
}
