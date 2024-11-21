package leetcode

import "slices"

func subsets(nums []int) [][]int {
	ans := [][]int{}
	parts := []int{}
	n := len(nums)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(parts))
			return
		}
		dfs(i + 1)
		parts = append(parts, nums[i])
		dfs(i + 1)
		parts = parts[:len(parts)-1]
	}
	dfs(0)
	return ans
}
