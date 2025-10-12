package leetcode

import "slices"

func threeSumTarget(nums []int, target int) [][]int {
	ans := [][]int{}
	n := len(nums)
	for i := range n - 2 {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		t := target - nums[i]
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

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	ans := [][]int{}
	n := len(nums)
	for a := range n - 3 {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for _, vals := range threeSumTarget(nums[a+1:], target-nums[a]) {
			ans = append(ans, append([]int{nums[a]}, vals...))
		}
	}
	return ans
}
