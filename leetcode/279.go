package leetcode

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		ans := math.MaxInt
		for j := 1; i >= j*j; j++ {
			ans = min(ans, 1+dp[i-j*j])
		}
		dp[i] = ans
	}
	return dp[n]
}
