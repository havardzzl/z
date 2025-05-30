package leetcode

import "math"

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][i] = 0
		for j := i + 1; j <= n; j++ {
			if j == i+1 {
				dp[i][j] = i
			} else {
				dp[i][j] = math.MaxInt
			}
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if dp[i][j] != math.MaxInt {
			return dp[i][j]
		}
		ans := math.MaxInt
		for k := i + 1; k < j; k++ {
			ans = min(ans, max(dfs(i, k-1), dfs(k+1, j))+k)
		}
		dp[i][j] = ans
		return ans
	}
	return dfs(1, n)
}
