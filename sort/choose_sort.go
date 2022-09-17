// @file: choose_sort.go
// @description: 选择排序
// @author: SaltFish
// @date: 2020/9/12

package sort

func chooseSort(nums []int) {
	for i := range nums {
		minIndex := i
		// 从未排序的部分选出最小的与排序部分的后一个数字交换
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}
