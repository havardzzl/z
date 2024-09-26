package leetcode

import "fmt"

func summaryRanges(nums []int) []string {
	n := len(nums)
	ans := []string{}
	if n == 0 {
		return ans
	}
	start, end := nums[0], nums[0]
	format := func(s, e int) string {
		if s == e {
			return fmt.Sprintf("%d", s)
		}
		return fmt.Sprintf("%d->%d", s, e)
	}
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] > 1 {
			ans = append(ans, format(start, nums[i-1]))
			start, end = nums[i], nums[i]
		} else {
			end = nums[i]
		}
	}
	ans = append(ans, format(start, end))
	return ans
}
