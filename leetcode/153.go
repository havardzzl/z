package leetcode

func findMin(nums []int) int {
	n := len(nums)
	left, right := 1, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < nums[0] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left < n {
		return nums[left]
	}
	return nums[0]
}
