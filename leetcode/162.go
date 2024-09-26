package leetcode

func findPeakElement(nums []int) int {
	i, j := 0, len(nums)-1
	for i+1 < j {
		h := int(uint(i+j) >> 1)
		if h == 0 || h == len(nums)-1 {
			return h
		} else if nums[h-1] < nums[h] && nums[h] < nums[h+1] {
			i = h
		} else if nums[h-1] < nums[h] && nums[h] > nums[h+1] {
			return h
		} else {
			j = h
		}
	}
	if nums[i] > nums[j] {
		return i
	}
	return j
}
