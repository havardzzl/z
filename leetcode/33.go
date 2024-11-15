package leetcode

func search2(nums []int, target int) int {
	n := len(nums)
	if nums[n-1] == target {
		return n - 1
	}
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if mid < n && nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[n-1] {
			if target < nums[mid] && target > nums[n-1] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target < nums[n-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if left > 0 && left < n && nums[left] == target {
		return left
	}
	return -1
}
