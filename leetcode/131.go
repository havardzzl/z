package leetcode

import (
	"slices"
)

func partition(s string) [][]string {
	ans := [][]string{}

	match := func(str string) bool {
		n := len(str)
		for i := 0; i < n/2; i++ {
			if str[i] != str[n-1-i] {
				return false
			}
		}
		return true
	}
	n := len(s)
	var dfs func(i int, parts []string)
	dfs = func(i int, parts []string) {
		if i == n {
			ans = append(ans, slices.Clone(parts))
			return
		}
		for j := i; j < n; j++ {
			if match(s[i : j+1]) {
				dfs(j+1, append(parts, s[i:j+1]))
			}
		}
	}
	dfs(0, []string{})
	return ans
}
