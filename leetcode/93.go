package leetcode

import (
	"strings"
)

func restoreIpAddresses(s string) []string {
	ans := []string{}
	parts := []string{}
	var dfs func(i int, cur string)
	dfs = func(i int, cur string) {
		if len(cur) == 3 && cur > "255" || len(cur) > 3 {
			return
		}
		if i < len(s) && len(parts) == 4 {
			return
		}
		if i == len(s) {
			ans = append(ans, strings.Join(parts, ""))
			return
		}
		newNum := s[i+1 : i+2]
		if cur == "0" {
			parts = append(parts, cur)
			dfs(i+1, newNum)
			parts = parts[:len(parts)-1]
			return
		}
		dfs(i+1, cur+newNum)
		parts = append(parts, cur)
		dfs(i+1, newNum)
		parts = parts[:len(parts)-1]
	}
	dfs(0, s[:1])
	return ans
}
