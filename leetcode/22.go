package leetcode

import "slices"

func generateParenthesis(n int) []string {
	ans := []string{}
	parts := []byte{}
	var dfs func(a, b int)
	dfs = func(a, b int) {
		if a == n && b == n {
			ans = append(ans, string(slices.Clone(parts)))
			return
		}
		if a == n {
			parts = append(parts, ')')
			dfs(a, b+1)
			parts = parts[:len(parts)-1]
			return
		}
		if b == a {
			parts = append(parts, '(')
			dfs(a+1, b)
			parts = parts[:len(parts)-1]
			return
		}
		parts = append(parts, '(')
		dfs(a+1, b)
		parts = parts[:len(parts)-1]
		parts = append(parts, ')')
		dfs(a, b+1)
		parts = parts[:len(parts)-1]
	}
	dfs(0, 0)
	return ans
}
