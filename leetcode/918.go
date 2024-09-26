package leetcode

func maxSubarraySumCircular(nums []int) int {
	nums2 := append(nums, nums...)

	dp := make([]int, len(nums2))
	dp[0] = nums2[0]

	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}
