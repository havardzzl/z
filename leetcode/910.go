package leetcode

import "sort"

func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)
	l := len(nums)
	ans := nums[l-1] - nums[0]
	for i := 0; i < l-1; i++ {
		ans = min(ans, max(nums[l-1]-k, nums[i]+k)-min(nums[i+1]-k, nums[0]+k))
	}
	return ans
}
