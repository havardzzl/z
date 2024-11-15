package leetcode

// search for min pos which value >= target
func searchLowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	left := searchLowerBound(nums, target)
	if left < n && nums[left] == target {
		right := searchLowerBound(nums, target+1) - 1
		return []int{left, right}
	}
	return []int{-1, -1}
}
