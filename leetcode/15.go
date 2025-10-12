package leetcode

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	ans := [][]int{}
	n := len(nums)
	for i := range n - 2 {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		t := -nums[i]
		for j < k {
			if nums[j]+nums[k] > t {
				k--
			} else if nums[j]+nums[k] < t {
				j++
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k--
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return ans
}
