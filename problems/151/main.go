package main

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)

	var res []string
	i, j := len(s)-1, len(s)-1
	for i >= 0 {
		for i >= 0 && s[i] != ' ' {
			i--
		}
		res = append(res, s[i+1:j+1])
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
	}
	return strings.Join(res, " ")
}
