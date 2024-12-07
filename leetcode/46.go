package leetcode

import "slices"

func permute(nums []int) [][]int {
	n := len(nums)
	visit := make([]bool, n)
	ans := [][]int{}
	parts := []int{}
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(parts))
			return
		}
		for j := 0; j < n; j++ {
			if visit[j] {
				continue
			}
			visit[j] = true
			parts = append(parts, nums[j])
			dfs(i + 1)
			parts = parts[:len(parts)-1]
			visit[j] = false
		}
	}
	dfs(0)
	return ans
}
