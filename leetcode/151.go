package leetcode

import (
	"strings"
)

func reverseWords(s string) string {
	ss := strings.Fields(s)
	n := len(ss)
	for i := 0; i < n/2; i++ {
		ss[i], ss[n-1-i] = ss[n-1-i], ss[i]
	}
	return strings.Join(ss, " ")
}
