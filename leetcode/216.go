package leetcode

import "slices"

func combinationSum3(k int, n int) [][]int {
	ans := [][]int{}
	nums := []int{}
	var sum int
	var dfs func(b int)
	dfs = func(b int) {
		d := k - len(nums)
		left := n - sum
		if d == 1 {
			if left >= b && left <= 9 {
				ans = append(ans, slices.Clone(append(nums, left)))
			}
			return
		}
		up := min(9, left)
		for i := b; i <= up; i++ {
			nums = append(nums, i)
			sum += i
			dfs(i + 1)
			sum -= i
			nums = nums[:len(nums)-1]
		}
	}
	dfs(1)
	return ans
}
