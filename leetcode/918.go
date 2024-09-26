package leetcode

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	dpMin := make([]int, n)
	dpMax := make([]int, n)
	dpMin[0], dpMax[0] = nums[0], nums[0]
	var sum int
	for _, num := range nums {
		sum += num
	}

	minDp := dpMin[0]
	maxDp := dpMax[0]

	for i := 1; i < len(nums); i++ {
		dpMin[i] = min(dpMin[i-1]+nums[i], nums[i])
		dpMax[i] = max(dpMax[i-1]+nums[i], nums[i])
		minDp = min(minDp, dpMin[i])
		maxDp = max(maxDp, dpMax[i])
	}
	if maxDp < 0 {
		return maxDp
	}
	return max(maxDp, sum-minDp)
}
