package leetcode

func numSubarrayProductLessThanK(nums []int, k int) int {
	product := 1
	ans := 0
	left := 0
	for i, v := range nums {
		if v >= k {
			product = 1
			left = i + 1
			continue
		}
		product *= v
		for product >= k {
			product /= nums[left]
			left++
		}
		ans += i - left + 1
	}
	return ans
}
