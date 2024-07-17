package main

func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

// palindrome 在 s 中寻找以 s[l] 和 s[r] 为中心的最长回文串
func palindrome(s string, l int, r int) string {
	for l >= 0 && r <= len(s)-1 && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
