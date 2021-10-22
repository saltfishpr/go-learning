// @file: 14-2剪绳子.go
// @date: 2021/2/13

// Package offer
package offer

/*
给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m - 1] 。请问 k[0]*k[1]*...*k[m - 1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
*/

func cuttingRope2X14(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b, p, x := n/3-1, n%3, 1000000007, 3

	pow := func(x, n int) int {
		if x == 0 {
			return 0
		}
		res := 1
		for n != 0 {
			// n是奇数
			if n&1 == 1 {
				res = (res * x) % p
			}
			x = (x * x) % p
			n >>= 1
		}
		return res
	}

	if b == 0 {
		return pow(x, a) * 3 % p
	}
	if b == 1 {
		return pow(x, a) * 4 % p
	}
	return pow(x, a) * 6 % p
}
