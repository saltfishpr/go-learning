// @file: 62圆圈中最后剩下的数字.go
// @date: 2021/2/24

// Package offer
package offer

/*
0,1,···,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。求出这个圆圈里剩下的最后一个数字。
例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。
*/

func lastRemainingX62(n int, m int) int {
	f := 0
	for i := 2; i < n+1; i++ {
		f = (m + f) % i
	}
	return f
}
