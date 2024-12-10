package leetcode

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	ans := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}
