package leetcode

import "slices"

func combine(n int, k int) [][]int {
	ans := [][]int{}
	nums := []int{}
	var dfs func(i int)
	dfs = func(i int) {
		cnt := len(nums)
		if cnt == k {
			ans = append(ans, slices.Clone(nums))
			return
		}
		if i < k-cnt {
			return
		}
		for j := i; j >= 1; j-- {
			nums = append(nums, j)
			dfs(j - 1)
			nums = nums[:len(nums)-1]
		}
	}
	dfs(n)
	return ans
}
