package leetcode

import (
	"slices"
	"sort"
)

func minimizeArrayValue(nums []int) int {
	return sort.Search(slices.Max(nums), func(limit int) bool {
		extra := 0
		for i := len(nums) - 1; i > 0; i-- {
			extra = max(extra+nums[i]-limit, 0)
		}
		return nums[0]+extra <= limit
	})
}
