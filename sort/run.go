// @file: run.go
// @description: 测试运行
// @author: SaltFish
// @date: 2020/9/11

// Package sort 实现排序算法
package sort

import "strconv"

func toString(nums []int) string {
	res := "["
	for _, n := range nums {
		res += strconv.Itoa(n) + ", "
	}
	res = res[:len(res)-2] + "]"
	return res
}

func RunSort() {
	nums := []int{12, 4, 132, 55, 46, 232, 789, 1, 0, 98, 523, 666}
	insertSort(nums)
	print("result: ", toString(nums), "\n")
}
