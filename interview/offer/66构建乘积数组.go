// @file: 66构建乘积数组.go
// @date: 2021/2/26

// Package offer
package offer

/*
给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
*/

func constructArrX66(a []int) []int {
	n := len(a)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = 1
	}
	for i := 1; i < n; i++ {
		b[i] = b[i-1] * a[i-1]
	}
	tmp := 1 // 记录右侧乘积
	for i := n - 2; i >= 0; i-- {
		tmp *= a[i+1]
		b[i] *= tmp
	}
	return b
}
