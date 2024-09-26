package leetcode

func minSubArrayLen(target int, nums []int) int {
	left, right, sum := 0, 0, 0
	minLen := len(nums) + 1
	for {
		if sum < target {
			if right == len(nums) {
				break
			}
			sum += nums[right]
			right++
		} else {
			for sum >= target {
				sum -= nums[left]
				left++
			}
			l := right - left + 1
			if l < minLen {
				minLen = l
			}
		}
	}
	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}
