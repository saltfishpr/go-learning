// @file: bubble_sort.go
// @description: 冒泡排序
// @author: SaltFish
// @date: 2020/9/11

package sort

func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		print("before: ", toString(nums), "\n")
		for j := 0; j < len(nums)-i-1; j++ { // 这里注意j的终止条件，已经排序好的最大的值已经在最后了
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		print("after:  ", toString(nums), "\n\n")
	}
}
