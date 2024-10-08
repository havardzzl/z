package leetcode

func canJump(nums []int) bool {
	trap := 0
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= trap {
			trap++
		} else {
			trap = 0
		}
	}
	return trap == 0
}
