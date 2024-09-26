package leetcode

func rob(nums []int) int {
	robNums := map[int]int{
		-2: 0,
		-1: 0,
		0:  nums[0],
	}
	for i := 1; i < len(nums); i++ {
		robNums[i] = max(robNums[i-3]+nums[i], robNums[i-2]+nums[i], robNums[i-1])
	}
	return robNums[len(nums)-1]
}
