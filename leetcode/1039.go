package leetcode

import "math"

func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if j-i == 1 {
			return 0
		}
		if dp[i][j] > 0 {
			return dp[i][j]
		}
		ans := math.MaxInt
		for k := i + 1; k < j; k++ {
			ans = min(ans, dfs(i, k)+dfs(k, j)+values[i]*values[j]*values[k])
		}
		dp[i][j] = ans
		return ans
	}
	return dfs(0, n-1)
}
