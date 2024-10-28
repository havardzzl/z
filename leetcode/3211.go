package leetcode

import (
	"strconv"
	"strings"
)

func validStrings(n int) []string {
	format := func(v int) string {
		ss := []string{}
		for i := n - 1; i >= 0; i-- {
			ss = append(ss, strconv.Itoa(v/(1<<i)))
			if v >= 1<<i {
				v -= 1 << i
			}
		}
		return strings.Join(ss, "")
	}
	ans := []string{}
	for i := 0; i < 1<<n; i++ {
		fs := format(i)
		if strings.Index(fs, "00") == -1 {
			ans = append(ans, fs)
		}
	}
	return ans
}
