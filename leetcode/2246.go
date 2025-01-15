package leetcode

import "slices"

func longestPath(parent []int, s string) int {
	sons := map[int][]int{}
	n := len(parent)
	for i := 1; i < n; i++ {
		if sons[parent[i]] == nil {
			sons[parent[i]] = []int{}
		}
		sons[parent[i]] = append(sons[parent[i]], i)
	}
	ans := 0
	var dfs func(k int) int
	dfs = func(k int) int {
		lengths := []int{}
		for _, son := range sons[k] {
			length := dfs(son)
			if s[k] != s[son] {
				lengths = append(lengths, length)
			}
		}
		slices.Sort(lengths)
		l := len(lengths)
		if l > 1 {
			ans = max(ans, lengths[l-1]+lengths[l-2]+1)
		} else if l > 0 {
			ans = max(ans, lengths[l-1]+1)
		} else {
			ans = max(ans, 1)
		}
		if l == 0 {
			return 1
		}
		return lengths[l-1] + 1
	}
	dfs(0)
	return ans
}
