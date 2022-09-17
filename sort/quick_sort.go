// @file: quick_sort.go
// @description: 快速排序
// @author: SaltFish
// @date: 2020/9/11

package sort

import "math/rand"

func myQuickSort(nums []int) {
	var partition func(int, int)
	partition = func(l, r int) {
		if l >= r {
			return
		}
		// print("before partition: ", toString(nums), "\n")
		i, j := l, r
		key := nums[l] // 将第一个数作为标准数
		for i < j {
			// 从后面找到第一个比key小的数
			for i < j && nums[j] >= key {
				j--
			}
			// 交换两个数并后移i
			if i < j {
				nums[i] = nums[j]
				i++
			}
			// 从前面找到第一个比key大的数
			for i < j && nums[i] < key {
				i++
			}
			// 交换两个数并前移j
			if i < j {
				nums[j] = nums[i]
				j--
			}
		}
		nums[i] = key // 将标准数填入中间
		// print("after partition:  ", toString(nums), "\n\n")
		partition(l, i-1) // 递归排序左侧数组
		partition(i+1, r) // 递归排序右侧数组
	}
	partition(0, len(nums)-1)
}

// 取最后一个数为key，遍历数组，将比key小的数移到左侧，大的数移到右侧，返回key所在位置
func partition(arr []int) (primeIdx int) {
	print("before partition: ", toString(arr), "\n")
	primeIdx = 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[len(arr)-1] {
			arr[i], arr[primeIdx] = arr[primeIdx], arr[i]
			primeIdx++
		}
	}
	arr[primeIdx], arr[len(arr)-1] = arr[len(arr)-1], arr[primeIdx]
	print("after partition:  ", toString(arr), "\n\n")
	return
}

func quickSort(arr []int) {
	if len(arr) > 1 {
		primeIdx := partition(arr)
		quickSort(arr[:primeIdx])
		quickSort(arr[primeIdx+1:])
	}
}

func randomQuickSort(arr []int) {
	if len(arr) > 1 {
		primeIdx := rand.Intn(len(arr))
		arr[primeIdx], arr[len(arr)-1] = arr[len(arr)-1], arr[primeIdx]
		primeIdx = partition(arr)
		randomQuickSort(arr[:primeIdx])
		randomQuickSort(arr[primeIdx+1:])
	}
}

func quickSortTail(arr []int) {
	for len(arr) > 1 {
		primeIdx := partition(arr)
		if primeIdx < len(arr)/2 {
			quickSortTail(arr[:primeIdx])
			arr = arr[primeIdx+1:]
		} else {
			quickSortTail(arr[primeIdx+1:])
			arr = arr[:primeIdx]
		}
	}
}
