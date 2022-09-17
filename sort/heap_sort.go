// @file: heap_sort.go
// @description: 堆排序
// @author: SaltFish
// @date: 2020/9/12

package sort

func left(x int) int {
	return x*2 + 1
}

func sink(i, n int, nums []int) {
	// 保存要下沉的节点值
	tmp := nums[i]
	for k := left(i); k < n; k = left(k) {
		// 如果有右孩子，则比较右孩子与左孩子的大小，取较大的那个
		if k+1 < n && nums[k] < nums[k+1] {
			k++
		}
		// 如果较大的孩子比要下沉的节点还要小，则不用继续下沉
		if nums[k] <= tmp {
			break
		}
		// 将较大的孩子交换上来
		nums[i] = nums[k]
		i = k
	}
	// 将要下沉的节点值保存在交换后的空位
	nums[i] = tmp
}

func heapSort(nums []int) {
	// 从最后一个有孩子的节点开始构建大顶堆
	for i := len(nums)/2 - 1; i >= 0; i-- {
		sink(i, len(nums), nums)
	}
	// 调整堆结构，交换堆顶元素与末尾元素
	for j := len(nums) - 1; j > 0; j-- {
		nums[0], nums[j] = nums[j], nums[0]
		sink(0, j, nums)
	}
}
