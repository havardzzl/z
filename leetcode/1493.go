package leetcode

func longestSubarray(nums []int) int {
	ans := 0
	curLen, lastLen := 0, 0
	var hasZero bool
	for _, num := range nums {
		if num == 1 {
			curLen++
		} else {
			ans = max(ans, curLen+lastLen)
			lastLen, curLen = curLen, 0
			hasZero = true
		}
	}
	if hasZero {
		ans = max(ans, curLen+lastLen)
	} else {
		ans = len(nums) - 1
	}
	return ans
}
