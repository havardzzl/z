package leetcode

import (
	"slices"
)

func triangleNumber(nums []int) int {
	slices.Sort(nums)
	n := len(nums)
	ans := 0
	for i := 2; i < n; i++ {
		j := 0
		k := i - 1
		for k > j {
			if nums[k]+nums[j] > nums[i] {
				ans += k - j
				k--
			} else {
				j++
			}
		}
	}
	return ans
}
