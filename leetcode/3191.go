package leetcode

func minOperations(nums []int) int {
	var step int
	rev := func(i int) int {
		return (i + 1) % 2
	}
	for len(nums) > 2 {
		if nums[0] == 0 {
			nums[1], nums[2] = rev(nums[1]), rev(nums[2])
			step++
		}
		nums = nums[1:]
	}
	allOne := true
	for _, num := range nums {
		if num != 1 {
			allOne = false
		}
	}
	if allOne {
		return step
	}
	return -1
}
