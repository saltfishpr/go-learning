// package sort
// @file: insert_sort.go
// @description: 插入排序
// @author: SaltFish
// @date: 2020/9/11
package sort

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		j := i - 1
		tmp := nums[i]
		// 找到对应的位置插入第i个数字
		for ; j >= 0; j-- {
			if tmp < nums[j] {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = tmp
	}
}
