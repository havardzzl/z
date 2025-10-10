package leetcode

func removeElement(nums []int, val int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		for r >= 0 && nums[r] == val {
			r--
		}
		for l < n && nums[l] != val {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	return l
}
