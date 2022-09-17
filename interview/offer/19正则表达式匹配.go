// @file: 19正则表达式匹配.go
// @date: 2021/2/14

// Package offer
package offer

/*
请实现一个函数用来匹配包含'.'和'*'的正则表达式。
模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。
在本题中，匹配是指字符串的所有字符匹配整个模式。
例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。

动态规划 dp[i][j] 代表字符串 s 的前 i 个字符和 p 的前 j 个字符能否匹配。
0 为空字符串，所以 dp[i][j] 对应的添加字符是 s[i - 1] 和 p[j - 1]
*/

func isMatchX19(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 2; i < n+1; i += 2 {
		dp[0][i] = dp[0][i-2] && p[i-1] == '*'
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if p[j-1] == '*' {
				// p[j-1]为'*'，则判断p[j-2]
				// p[j-2]出现0次 || p[j-2]出现1次 || p[j-2]出现2次 || p[j-2] 为任意字符 '.'
				dp[i][j] = dp[i][j-2] || dp[i][j-1] || dp[i-1][j] && s[i-1] == p[j-2] ||
					dp[i-1][j] && p[j-2] == '.'
			} else {
				dp[i][j] = dp[i-1][j-1] && s[i-1] == p[j-1] || dp[i-1][j-1] && p[j-1] == '.'
			}
		}
	}
	return dp[m][n]
}
