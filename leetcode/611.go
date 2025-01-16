package leetcode

import (
	"slices"
	"sort"
)

func triangleNumber(nums []int) int {
	slices.Sort(nums)
	n := len(nums)
	ans := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			maxV := nums[i] + nums[j]
			idx := sort.SearchInts(nums, maxV)
			if idx > j+1 {
				ans += idx - j - 1
			}
		}
	}
	return ans
}
