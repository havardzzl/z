package leetcode

func findMin2(nums []int) int {
	n := len(nums)
	l := 0
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[n-1] {
			l++
		} else {
			break
		}
	}
	nums = nums[l:]
	n = len(nums)
	r := 0
	for i := n - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			r++
		} else {
			break
		}
	}
	nums = nums[:n-r]
	l, r = 0, len(nums)-2
	n = len(nums)
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] <= nums[n-1] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return min(nums[l], nums[n-1])
}
