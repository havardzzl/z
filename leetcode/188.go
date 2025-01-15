package leetcode

import (
	"math"
)

func maxProfit3(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, k+2)
	for i := 0; i < k+2; i++ {
		dp[i][1] = -prices[0]
	}
	dp[0][1] = math.MinInt16
	for i := 1; i < n; i++ {
		for j := 1; j < k+2; j++ {
			newJ0 := max(dp[j][0], dp[j-1][1]+prices[i])
			dp[j][1] = max(dp[j][1], dp[j][0]-prices[i])
			dp[j][0] = newJ0
		}
	}
	return dp[k+1][0]
}
