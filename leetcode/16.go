package leetcode

import (
	"slices"
)

func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)
	n := len(nums)
	ans := nums[0] + nums[1] + nums[2]
	var dist = func(val int) int {
		if val > target {
			return val - target
		}
		return target - val
	}
	for i := range n - 2 {
		j, k := i+1, n-1
		for j < k {
			total := nums[i] + nums[j] + nums[k]
			if dist(ans) > dist(total) {
				ans = total
			}
			if total == target {
				return target
			} else if total < target {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}
